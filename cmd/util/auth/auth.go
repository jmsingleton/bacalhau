package auth

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/bacalhau-project/bacalhau/cmd/util/choose"
	"github.com/bacalhau-project/bacalhau/pkg/authn"
	"github.com/bacalhau-project/bacalhau/pkg/authn/ask"
	"github.com/bacalhau-project/bacalhau/pkg/authn/challenge"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/publicapi/apimodels"
	"github.com/bacalhau-project/bacalhau/pkg/publicapi/client/v2"
	"github.com/bacalhau-project/bacalhau/pkg/system"
)

type Responder interface {
	Respond(request *json.RawMessage) ([]byte, error)
}

func RunAuthenticationFlow(cmd *cobra.Command, auth *client.Auth, c config.Context) (*apimodels.HTTPCredential, error) {
	sk, err := config.GetClientPrivateKey(c)
	if err != nil {
		return nil, err
	}
	supportedMethods := map[authn.MethodType]Responder{
		authn.MethodTypeChallenge: &challenge.Responder{Config: c, Signer: system.NewMessageSigner(sk)},
		authn.MethodTypeAsk:       &ask.Responder{Cmd: cmd},
	}

	methods, err := auth.Methods(cmd.Context(), &apimodels.ListAuthnMethodsRequest{})
	if err != nil {
		return nil, err
	}

	filteredMethods := make(map[string]authn.Requirement, len(methods.Methods))
	clientTypes := maps.Keys(supportedMethods)
	for name, req := range methods.Methods {
		if lo.Contains(clientTypes, req.Type) {
			filteredMethods[name] = req
		}
	}

	if len(filteredMethods) == 0 {
		serverTypes := lo.Map(maps.Values(methods.Methods), func(r authn.Requirement, _ int) authn.MethodType { return r.Type })
		return nil, fmt.Errorf("no common authentication method: client supports %v, server supports %v", clientTypes, serverTypes)
	}

	var authentication authn.Authentication
	for !authentication.Success {
		supportedNames := maps.Keys(filteredMethods)
		chosenMethodName, err := choose.Choose(cmd, "How would you like to authenticate?", supportedNames)
		if errors.Is(err, io.EOF) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		methodRequirement := methods.Methods[chosenMethodName]
		methodResponder := supportedMethods[methodRequirement.Type]
		response, err := methodResponder.Respond(methodRequirement.Params)
		if err != nil {
			return nil, err
		}

		authnResponse, err := auth.Authenticate(cmd.Context(), &apimodels.AuthnRequest{
			Name:       chosenMethodName,
			MethodData: response,
		})
		if err != nil {
			return nil, err
		}

		authentication = authnResponse.Authentication
		if authentication.Reason != "" {
			cmd.PrintErrln(authentication.Reason)
		}
	}

	return &apimodels.HTTPCredential{
		Scheme: "Bearer",
		Value:  authentication.Token,
	}, nil
}

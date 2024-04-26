package util

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/models"
	"github.com/bacalhau-project/bacalhau/pkg/publicapi/client"
	"github.com/bacalhau-project/bacalhau/pkg/system"

	"github.com/bacalhau-project/bacalhau/pkg/version"
)

type Versions struct {
	ClientVersion *models.BuildVersionInfo `json:"clientVersion,omitempty"`
	ServerVersion *models.BuildVersionInfo `json:"serverVersion,omitempty"`
	LatestVersion *models.BuildVersionInfo `json:"latestVersion,omitempty"`
	UpdateMessage string                   `json:"updateMessage,omitempty"`
}

func GetAllVersions(ctx context.Context, c config.Context) (Versions, error) {
	var err error
	versions := Versions{ClientVersion: version.Get()}

	legacyTLS := client.LegacyTLSSupport(config.ClientTLSConfig(c))
	sk, err := config.GetClientPrivateKey(c)
	if err != nil {
		return versions, err
	}
	signer := system.NewMessageSigner(sk)
	client, err := client.NewAPIClient(legacyTLS, c, signer, config.ClientAPIHost(c), config.ClientAPIPort(c))
	if err != nil {
		return versions, err
	}
	versions.ServerVersion, err = client.Version(ctx)
	if err != nil {
		return versions, errors.Wrap(err, "error running version command")
	}

	clientID, err := config.GetClientID(c)
	if err != nil {
		return versions, errors.Wrap(err, "error getting client ID")
	}

	InstallationID, err := config.GetInstallationUserID(c)
	if err != nil {
		return versions, errors.Wrap(err, "error getting Installation ID")
	}

	updateCheck, err := version.CheckForUpdate(ctx, versions.ClientVersion, versions.ServerVersion, clientID, InstallationID)
	if err != nil {
		return versions, errors.Wrap(err, "failed to get latest version")
	} else {
		versions.UpdateMessage = updateCheck.Message
		versions.LatestVersion = updateCheck.Version
	}

	return versions, nil
}

func EnsureValidVersion(ctx context.Context, clientVersion, serverVersion *models.BuildVersionInfo) error {
	if clientVersion == nil {
		log.Ctx(ctx).Warn().Msg("Unable to parse nil client version, skipping version check")
		return nil
	}
	if clientVersion.GitVersion == version.DevelopmentGitVersion {
		log.Ctx(ctx).Debug().Msg("Development client version, skipping version check")
		return nil
	}
	if serverVersion == nil {
		log.Ctx(ctx).Warn().Msg("Unable to parse nil server version, skipping version check")
		return nil
	}
	if serverVersion.GitVersion == version.DevelopmentGitVersion {
		log.Ctx(ctx).Debug().Msg("Development server version, skipping version check")
		return nil
	}
	c, err := semver.NewVersion(clientVersion.GitVersion)
	if err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("Unable to parse client version, skipping version check")
		return nil
	}
	s, err := semver.NewVersion(serverVersion.GitVersion)
	if err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("Unable to parse server version, skipping version check")
		return nil
	}
	if s.GreaterThan(c) {
		return fmt.Errorf(`the server version %s is newer than client version %s, please upgrade your client with the following command:
curl -sL https://get.bacalhau.org/install.sh | bash`,
			serverVersion.GitVersion,
			clientVersion.GitVersion,
		)
	}
	if c.GreaterThan(s) {
		return fmt.Errorf(
			"client version %s is newer than server version %s, please ask your network administrator to update Bacalhau",
			clientVersion.GitVersion,
			serverVersion.GitVersion,
		)
	}
	return nil
}

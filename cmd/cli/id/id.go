package id

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/bacalhau-project/bacalhau/cmd/util"
	"github.com/bacalhau-project/bacalhau/cmd/util/flags/cliflags"
	"github.com/bacalhau-project/bacalhau/cmd/util/flags/configflags"
	"github.com/bacalhau-project/bacalhau/cmd/util/output"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	bac_libp2p "github.com/bacalhau-project/bacalhau/pkg/libp2p"
	"github.com/bacalhau-project/bacalhau/pkg/util/closer"
)

type IDInfo struct {
	ID       string `json:"ID"`
	ClientID string `json:"ClientID"`
}

func NewCmd(cfg config.Context) *cobra.Command {
	outputOpts := output.OutputOptions{
		Format: output.JSONFormat,
	}

	idFlags := map[string][]configflags.Definition{
		"libp2p": configflags.Libp2pFlags,
	}

	idCmd := &cobra.Command{
		Use:   "id",
		Short: "Show bacalhau node id info",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return configflags.BindFlags(cmd, cfg, idFlags)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return id(cmd, cfg, outputOpts)
		},
	}

	// TODO(forrest): [ux] these are flags without a corresponding value in the config
	// in the future we can bind all flags to a config value.
	idCmd.Flags().AddFlagSet(cliflags.OutputFormatFlags(&outputOpts))

	if err := configflags.RegisterFlags(idCmd, idFlags); err != nil {
		util.Fatal(idCmd, err, 1)
	}

	return idCmd
}

var idColumns = []output.TableColumn[IDInfo]{
	{
		ColumnConfig: table.ColumnConfig{Name: "id"},
		Value:        func(i IDInfo) string { return i.ID },
	},
	{
		ColumnConfig: table.ColumnConfig{Name: "client id"},
		Value:        func(i IDInfo) string { return i.ClientID },
	},
}

func id(cmd *cobra.Command, cfg config.Context, outputOpts output.OutputOptions) error {
	privKey, err := config.GetLibp2pPrivKey(cfg)
	if err != nil {
		return err
	}
	libp2pCfg, err := config.GetLibp2pConfig(cfg)
	if err != nil {
		return err
	}

	libp2pHost, err := bac_libp2p.NewHost(libp2pCfg.SwarmPort, privKey)
	if err != nil {
		return err
	}
	defer closer.CloseWithLogOnError("libp2pHost", libp2pHost)

	clientID, err := config.GetClientID(cfg)
	if err != nil {
		return err
	}
	info := IDInfo{
		ID:       libp2pHost.ID().String(),
		ClientID: clientID,
	}

	return output.OutputOne(cmd, idColumns, outputOpts, info)
}

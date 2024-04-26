package config

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/bacalhau-project/bacalhau/cmd/util/flags/cliflags"
	"github.com/bacalhau-project/bacalhau/cmd/util/hook"
	"github.com/bacalhau-project/bacalhau/cmd/util/output"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"
)

func newListCmd(cfg config.Context) *cobra.Command {
	o := output.OutputOptions{
		Format:     output.TableFormat,
		Pretty:     true,
		HideHeader: false,
		NoStyle:    false,
		Wide:       false,
	}
	listCmd := &cobra.Command{
		Use:      "list",
		Short:    "List all config keys.",
		Args:     cobra.MinimumNArgs(0),
		PreRunE:  hook.ClientPreRunHooks,
		PostRunE: hook.ClientPostRunHooks,
		RunE: func(cmd *cobra.Command, args []string) error {
			// create or open a repo
			repoPath, err := cfg.RepoPath()
			if err != nil {
				return err
			}
			_, err = setup.SetupBacalhauRepo(repoPath, cfg)
			if err != nil {
				return err
			}
			return list(cmd, cfg, o)
		},
	}
	listCmd.Flags().AddFlagSet(cliflags.OutputFormatFlags(&o))
	return listCmd
}

type configListEntry struct {
	Key   string
	Value interface{}
}

func list(cmd *cobra.Command, cfg config.Context, o output.OutputOptions) error {
	o.SortBy = []table.SortBy{{
		Name: "Key",
		Mode: table.Asc,
	}}
	var cfgList []configListEntry
	for _, k := range cfg.User().AllKeys() {
		v := cfg.User().Get(k)
		cfgList = append(cfgList, configListEntry{
			Key:   k,
			Value: v,
		})
	}

	if err := output.Output(cmd, listColumns, o, cfgList); err != nil {
		return err
	}

	return nil
}

var listColumns = []output.TableColumn[configListEntry]{
	{
		ColumnConfig: table.ColumnConfig{Name: "Key"},
		Value: func(s configListEntry) string {
			return s.Key
		},
	},
	{
		ColumnConfig: table.ColumnConfig{Name: "Value", WidthMax: 40, WidthMaxEnforcer: text.WrapHard},
		Value: func(v configListEntry) string {
			return fmt.Sprintf("%v", v.Value)
		},
	},
}

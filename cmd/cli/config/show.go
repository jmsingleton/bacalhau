package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"
)

func newShowCmd(cfg config.Context) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Show the current bacalhau config.",
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

			return showConfig(cmd, cfg)
		},
	}
	showCmd.PersistentFlags().String("path", viper.GetString("repo"), "sets path dependent config fields")
	return showCmd
}

func showConfig(cmd *cobra.Command, cfg config.Context) error {
	// clear any existing configuration before generating the current.
	currentConfig, err := cfg.Current()
	if err != nil {
		return err
	}
	cfgbytes, err := yaml.Marshal(currentConfig)
	if err != nil {
		return err
	}
	cmd.Println(string(cfgbytes))
	return nil
}

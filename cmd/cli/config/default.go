package config

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/bacalhau-project/bacalhau/pkg/config"
)

func newDefaultCmd(cfg config.Context) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "default",
		Short: "Show the default bacalhau config.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return defaultConfig(cmd)
		},
	}
	showCmd.PersistentFlags().String("path", cfg.System().GetString("repo"), "sets path dependent config fields")
	return showCmd
}

func defaultConfig(cmd *cobra.Command) error {
	// clear any existing configuration before generating the default.
	c := config.New()
	defaultConfig, err := c.Current()
	if err != nil {
		return err
	}
	cfgbytes, err := yaml.Marshal(defaultConfig)
	if err != nil {
		return err
	}
	cmd.Println(string(cfgbytes))
	return nil
}

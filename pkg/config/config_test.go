//go:build unit || !integration

package config_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/config/configenv"
	"github.com/bacalhau-project/bacalhau/pkg/config/types"
	"github.com/bacalhau-project/bacalhau/pkg/repo"
)

func TestConfig(t *testing.T) {
	// Cleanup viper settings after each test

	// Testing Set and Get
	t.Run("SetAndGetHappyPath", func(t *testing.T) {
		expectedConfig := configenv.Testing
		c := config.New(config.WithDefaultConfig(expectedConfig))

		var out types.NodeConfig
		err := c.ForKey(types.Node, &out)
		assert.NoError(t, err)
		assert.Equal(t, expectedConfig.Node, out)

		retrieved, err := config.Get[string](c, types.NodeServerAPIHost)
		assert.NoError(t, err)
		assert.Equal(t, expectedConfig.Node.ServerAPI.Host, retrieved)
	})
	t.Run("SetAndGetAdvance", func(t *testing.T) {
		expectedConfig := configenv.Testing
		expectedConfig.Node.IPFS.SwarmAddresses = []string{"1", "2", "3", "4", "5"}
		c := config.New(config.WithDefaultConfig(expectedConfig))

		var out types.IpfsConfig
		err := c.ForKey(types.NodeIPFS, &out)
		assert.NoError(t, err)
		assert.Equal(t, expectedConfig.Node.IPFS, out)

		var node types.NodeConfig
		err = c.ForKey(types.Node, &node)
		assert.Equal(t, expectedConfig.Node, node)
		assert.NoError(t, err)

		var invalidNode types.NodeConfig
		err = c.ForKey("INVALID", &invalidNode)
		assert.Error(t, err)
	})

	// Testing KeyAsEnvVar
	t.Run("KeyAsEnvVar", func(t *testing.T) {
		assert.Equal(t, "BACALHAU_NODE_SERVERAPI_HOST", config.KeyAsEnvVar(types.NodeServerAPIHost))
	})

	// Testing Init
	t.Run("Init", func(t *testing.T) {
		testCases := []struct {
			name       string
			configType string
		}{
			{"config", "yaml"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				expected := configenv.Testing
				c := config.New(config.WithDefaultConfig(expected))

				var out types.NodeConfig
				err := c.ForKey(types.Node, &out)
				assert.NoError(t, err)
				assert.Equal(t, expected.Node.Requester, out.Requester)

				var retrieved types.RequesterConfig
				require.NoError(t, c.ForKey(types.NodeRequester, &retrieved))
				assert.Equal(t, expected.Node.Requester, retrieved)
			})
		}

	})

	t.Run("Load", func(t *testing.T) {
		testCases := []struct {
			name       string
			configType string
		}{
			{"yaml config type", "yaml"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// TODO(forrest): we are skipping this test because at present the repo nor config package
				// actually write the config, and it instead happens via a one-off method in NodeConstruction
				t.Skip()
				// define a configuration, init a repo with it, and assert the config was loaded from the repo.
				expected := configenv.Testing
				configPath := t.TempDir()
				expected.Node.Requester.JobStore.Path = filepath.Join(configPath, config.OrchestratorJobStorePath)
				c := config.New(config.WithDefaultConfig(expected))

				r, err := repo.NewFS(repo.FsRepoParams{
					Path:       configPath,
					Migrations: nil,
				})
				require.NoError(t, err)
				err = r.Init(c)
				require.NoError(t, err)

				// Now, try to load the configuration we just saved.
				err = c.Load(filepath.Join(configPath, config.FileName))
				require.NoError(t, err)
				loadedConfig, err := c.Current()
				require.NoError(t, err)

				// After loading, compare the loaded configuration with the expected configuration.
				assert.Equal(t, expected.Node.Requester, loadedConfig.Node.Requester)

				// Further, test specific parts:
				var out types.NodeConfig
				err = c.ForKey(types.Node, &out)
				assert.NoError(t, err)
				assert.Equal(t, expected.Node.Requester, out.Requester)

				var retrieved types.RequesterConfig
				require.NoError(t, c.ForKey(types.NodeRequester, &retrieved))
				assert.Equal(t, expected.Node.Requester, retrieved)
			})
		}
	})
}

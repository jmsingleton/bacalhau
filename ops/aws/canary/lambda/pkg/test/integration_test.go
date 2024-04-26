//go:build integration || !unit

package test

import (
	"context"
	"os"
	"testing"

	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/models"
	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/router"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"

	"github.com/stretchr/testify/require"
)

func TestScenariosAgainstProduction(t *testing.T) {
	// init system configs and repo.
	repoPath, err := os.MkdirTemp("", "bacalhau_canary_repo_*")
	require.NoError(t, err)
	cfg := config.New()
	_, err = setup.SetupBacalhauRepo(repoPath, cfg)
	require.NoError(t, err)

	for name := range router.TestcasesMap {
		t.Run(name, func(t *testing.T) {
			if name == "submitDockerIPFSJobAndGet" {
				t.Skip("skipping submitDockerIPFSJobAndGet as it is not stable yet. " +
					"https://github.com/bacalhau-project/bacalhau/issues/1869")
				return
			}
			event := models.Event{Action: name}
			err := router.Route(context.Background(), cfg, event)
			require.NoError(t, err)
		})
	}
}

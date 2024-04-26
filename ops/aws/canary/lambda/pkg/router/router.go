package router

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/models"
	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/scenarios"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"
)

var TestcasesMap = map[string]Handler{
	"list":                      scenarios.List,
	"submit":                    scenarios.Submit,
	"submitAndGet":              scenarios.SubmitAndGet,
	"submitDockerIPFSJobAndGet": scenarios.SubmitDockerIPFSJobAndGet,
	"submitAndDescribe":         scenarios.SubmitAnDescribe,
	"submitWithConcurrency":     scenarios.SubmitWithConcurrency,
}

var cfg config.Context

func init() {
	repoPath, err := os.MkdirTemp("", "bacalhau_canary_repo_*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create repo dir: %s", err)
		os.Exit(1)
	}

	cfg = config.New()
	// init system configs and repo.
	_, err = setup.SetupBacalhauRepo(repoPath, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize bacalhau repo: %s", err)
		os.Exit(1)
	}
}

func Route(ctx context.Context, cfg config.Context, event models.Event) error {
	handler, ok := TestcasesMap[event.Action]
	if !ok {
		return fmt.Errorf("no handler found for action: %s", event.Action)
	}
	err := handler(ctx, cfg)
	if err != nil {
		return fmt.Errorf("testcase %s failed: %s", event.Action, err)
	}
	log.Info().Msgf("testcase %s passed", event.Action)
	return nil
}

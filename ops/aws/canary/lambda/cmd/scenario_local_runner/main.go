package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/models"
	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/router"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"

	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func main() {
	var action string
	flag.StringVar(&action, "action", "",
		"Action to test. Useful when testing locally before pushing to lambda")
	flag.Parse()

	log.Info().Msgf("Testing locally the action: %s", action)

	// init system configs and repo.
	repoPath, err := os.MkdirTemp("", "bacalhau_canary_repo_*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create repo dir: %s", err)
		os.Exit(1)
	}

	cfg := config.New()
	_, err = setup.SetupBacalhauRepo(repoPath, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize bacalhau repo: %s", err)
		os.Exit(1)
	}

	err = router.Route(context.Background(), cfg, models.Event{Action: action})
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/logger"
	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/router"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"
)

func init() {
	logger.SetupCWLogger()
}

func main() {
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

	// running in lambda
	lambda.Start(router.Route)
}

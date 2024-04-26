package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/models"
	"github.com/bacalhau-project/bacalhau/ops/aws/canary/pkg/router"
	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/setup"

	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func main() {
	// parse flags
	var rate float32
	flag.Float32Var(&rate, "rate", 1.0,
		"Rate to execute each scenario. e.g. 0.1 means 1 execution every 10 seconds for each scenario")

	flag.Parse()
	log.Info().Msgf("Starting canary with rate: %f ", rate)

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

	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(1)
	}()

	for action := range router.TestcasesMap {
		go run(ctx, cfg, action, rate)
	}

	<-ctx.Done()
}

func run(ctx context.Context, cfg config.Context, action string, rate float32) {
	log.Ctx(ctx).Info().Msgf("Starting scenario: %s", action)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := router.Route(ctx, cfg, models.Event{Action: action})
			if err != nil {
				log.Ctx(ctx).Error().Msg(err.Error())
			}
		}
		jitter := rand.Intn(100) - 100 // +- 100ms sleep jitter
		time.Sleep(time.Duration(1/rate)*time.Second + time.Duration(jitter)*time.Millisecond)
	}
}

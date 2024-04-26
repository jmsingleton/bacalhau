package scenarios

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/bacalhau/pkg/config"
)

func Submit(ctx context.Context, cfg config.Context) error {
	// intentionally delay creation of the client so a new client is created for each
	// scenario to mimic the behavior of bacalhau cli.
	client, err := getClient(cfg)
	if err != nil {
		return err
	}

	j, err := getSampleDockerJob()
	if err != nil {
		return err
	}
	submittedJob, err := client.Submit(ctx, j)
	if err != nil {
		return err
	}

	log.Info().Msgf("submitted job: %s", submittedJob.Metadata.ID)

	err = waitUntilCompleted(ctx, client, submittedJob)
	if err != nil {
		return err
	}
	return nil
}

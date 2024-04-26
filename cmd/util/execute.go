package util

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	legacy_job "github.com/bacalhau-project/bacalhau/pkg/legacyjob"
	"github.com/bacalhau-project/bacalhau/pkg/model"
)

//nolint:funlen,gocyclo // Refactor later
func ExecuteJob(ctx context.Context, j *model.Job, c config.Context) (*model.Job, error) {
	err := legacy_job.VerifyJob(ctx, j)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Job failed to validate.")
		return nil, err
	}

	api, err := GetAPIClient(c)
	if err != nil {
		return nil, err
	}
	return api.Submit(ctx, j)
}

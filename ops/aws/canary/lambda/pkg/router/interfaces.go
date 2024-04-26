package router

import (
	"context"

	"github.com/bacalhau-project/bacalhau/pkg/config"
)

type Handler func(ctx context.Context, cfg config.Context) error

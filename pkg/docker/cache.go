package docker

import (
	"github.com/bacalhau-project/bacalhau/pkg/cache"
	"github.com/bacalhau-project/bacalhau/pkg/cache/basic"
	"github.com/bacalhau-project/bacalhau/pkg/config"
)

func NewManifestCache(cfg config.Context) cache.Cache[ImageManifest] {
	settings, _ := config.GetDockerManifestCacheSettings(cfg)

	// Used by compute nodes to map requester provided image identifiers (with
	// digest) to
	c, _ := basic.NewCache[ImageManifest](
		basic.WithCleanupFrequency(settings.Frequency.AsTimeDuration()),
		basic.WithMaxCost(settings.Size),
		basic.WithTTL(settings.Duration.AsTimeDuration()),
	)
	return c
}

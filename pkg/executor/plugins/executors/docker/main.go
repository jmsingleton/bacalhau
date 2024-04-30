package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/executor/docker"
	"github.com/bacalhau-project/bacalhau/pkg/executor/plugins/grpc"
)

const PluggableExecutorPluginName = "PLUGGABLE_EXECUTOR"

// HandshakeConfig is used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "EXECUTOR_PLUGIN",
	MagicCookieValue: "bacalhau_executor",
}

func main() { // Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "docker-plugin",
		Output: os.Stderr,
		Level:  hclog.Trace,
	})
	c := config.New()
	dockerExecutor, err := docker.NewExecutor(
		"bacalhau-pluggable-executor-docker",
		c,
	)
	if err != nil {
		logger.Error(err.Error())
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			PluggableExecutorPluginName: &grpc.ExecutorGRPCPlugin{
				Impl: dockerExecutor,
			},
		},
		Logger:     logger,
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

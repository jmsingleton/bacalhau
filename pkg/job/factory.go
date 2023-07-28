package job

import (
	"fmt"
	"time"

	"github.com/bacalhau-project/bacalhau/pkg/model"
	"github.com/bacalhau-project/bacalhau/pkg/system"
	"golang.org/x/exp/maps"
)

type SpecOpt func(s *model.Spec) error

func WithPublisher(p model.PublisherSpec) SpecOpt {
	return func(s *model.Spec) error {
		s.Publisher = p.Type
		s.PublisherSpec = p
		return nil
	}
}

func WithNetwork(network model.Network, domains []string) SpecOpt {
	return func(s *model.Spec) error {
		s.Network.Type = network
		s.Network.Domains = domains
		return nil
	}
}

func WithResources(cpu, memory, disk, gpu string) SpecOpt {
	return func(s *model.Spec) error {
		s.Resources.CPU = cpu
		s.Resources.Memory = memory
		s.Resources.Disk = disk
		s.Resources.GPU = gpu
		return nil
	}
}

func WithTimeout(t float64) SpecOpt {
	return func(s *model.Spec) error {
		s.Timeout = t
		return nil
	}
}

func WithDeal(targeting model.TargetingMode, concurrency int) SpecOpt {
	return func(s *model.Spec) error {
		s.Deal.TargetingMode = targeting
		s.Deal.Concurrency = concurrency
		return nil
	}
}

func WithAnnotations(annotations ...string) SpecOpt {
	return func(s *model.Spec) error {
		s.Annotations = annotations
		return nil
	}
}

func WithInputs(inputs ...model.StorageSpec) SpecOpt {
	return func(s *model.Spec) error {
		s.Inputs = inputs
		return nil
	}
}

func WithOutputs(outputs ...model.StorageSpec) SpecOpt {
	return func(s *model.Spec) error {
		s.Outputs = outputs
		return nil
	}
}

func WithNodeSelector(selector []model.LabelSelectorRequirement) SpecOpt {
	return func(s *model.Spec) error {
		s.NodeSelectors = selector
		return nil
	}
}

func WithDockerEngine(image, workdir string, entrypoint, parameters []string) SpecOpt {
	return func(s *model.Spec) error {
		if err := system.ValidateWorkingDir(workdir); err != nil {
			return fmt.Errorf("validating docker working directory: %w", err)
		}
		s.Engine = model.EngineDocker
		s.Docker = model.JobSpecDocker{
			Image:            image,
			Entrypoint:       entrypoint,
			Parameters:       parameters,
			WorkingDirectory: workdir,
		}
		return nil
	}
}

func WithEnvironmentVariables(vars map[string]string) SpecOpt {
	return func(s *model.Spec) error {
		maps.Copy(s.EnvironmentVariables, vars)
		return nil
	}
}

func MakeDockerSpec(
	image, workingdir string, entrypoint, parameters []string,
	opts ...SpecOpt,
) (model.Spec, error) {
	spec, err := MakeSpec(append(opts, WithDockerEngine(image, workingdir, entrypoint, parameters))...)
	if err != nil {
		return model.Spec{}, err
	}
	return spec, nil
}

func WithWasmEngine(
	entryModule model.StorageSpec,
	entrypoint string,
	parameters []string,
	importModules []model.StorageSpec,
) SpecOpt {
	return func(s *model.Spec) error {

		s.Engine = model.EngineWasm
		s.Wasm = model.JobSpecWasm{
			EntryModule:   entryModule,
			EntryPoint:    entrypoint,
			Parameters:    parameters,
			ImportModules: importModules,
		}
		return nil
	}
}
func MakeWasmSpec(
	entryModule model.StorageSpec, entrypoint string, parameters []string, importModules []model.StorageSpec,
	opts ...SpecOpt,
) (model.Spec, error) {
	spec, err := MakeSpec(append(opts, WithWasmEngine(entryModule, entrypoint, parameters, importModules))...)
	if err != nil {
		return model.Spec{}, err
	}
	return spec, nil
}

// TODO(forrest): find a home
const DefaultTimeout = 30 * time.Minute

func MakeSpec(opts ...SpecOpt) (model.Spec, error) {
	spec := &model.Spec{
		Engine:    model.EngineNoop,
		Publisher: model.PublisherNoop,
		PublisherSpec: model.PublisherSpec{
			Type: model.PublisherNoop,
		},
		Resources: model.ResourceUsageConfig{},
		Network: model.NetworkConfig{
			Type: model.NetworkNone,
		},
		Timeout: float64(DefaultTimeout),
		Deal: model.Deal{
			Concurrency: 1,
		},
		EnvironmentVariables: make(map[string]string),
	}

	for _, opt := range opts {
		if err := opt(spec); err != nil {
			return model.Spec{}, err
		}
	}

	return *spec, nil
}

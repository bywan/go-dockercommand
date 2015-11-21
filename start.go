package dockercommand

import (
	docker "github.com/fsouza/go-dockerclient"
)

type StartOptions struct {
	ID              string
	VolumeBinds     []string
	Links           []string
	PublishAllPorts bool
	PortBindings    map[docker.Port][]docker.PortBinding
	NetworkMode     string
	LogConfig       docker.LogConfig
}

func (dock *Docker) Start(options *StartOptions) error {
	return dock.client.StartContainer(options.ID, &docker.HostConfig{
		Binds:           options.VolumeBinds,
		Links:           options.Links,
		PublishAllPorts: options.PublishAllPorts,
		PortBindings:    options.PortBindings,
		NetworkMode:     options.NetworkMode,
		LogConfig:       options.LogConfig,
	})
}

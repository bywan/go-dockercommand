package dockercommand

import (
	docker "github.com/fsouza/go-dockerclient"
)

type StartOptions struct {
	ID string
}

func (dock *Docker) Start(options *StartOptions) error {
	return dock.client.StartContainer(options.ID, &docker.HostConfig{})
}

package dockercommand

import (
	"bufio"
	"fmt"
	"io"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

type RunOptions struct {
	Image       string
	Cmd         []string
	VolumeBinds []string
	Detach      bool
	Env         map[string]string
}

func (dock *Docker) Run(options *RunOptions) (*Container, error) {
	if err := dock.pullImageIfNotExist(options.Image); err != nil {
		return nil, err
	}

	container, err := dock.client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: options.Image,
			Cmd:   options.Cmd,
			Env:   convertEnvMapToSlice(options.Env),
		},
	})
	if err != nil {
		return nil, err
	}

	err = dock.client.StartContainer(container.ID, &docker.HostConfig{
		Binds: options.VolumeBinds,
	})
	if err != nil {
		return nil, err
	}

	containerWrapper := &Container{
		info:   container,
		client: dock.client,
	}

	if !options.Detach {
		_, err = dock.client.WaitContainer(container.ID)
		if err != nil {
			return nil, err
		}
	}
	return containerWrapper, nil
}

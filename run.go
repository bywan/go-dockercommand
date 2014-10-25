package dockercommand

import docker "github.com/fsouza/go-dockerclient"

type RunOptions struct {
	Image       string
	Cmd         []string
	VolumeBinds []string
	Detach      bool
	Env         map[string]string
}

func (dock *Docker) Run(options *RunOptions) (string, error) {
	if err := dock.pullImageIfNotExist(options.Image); err != nil {
		return "", err
	}

	container, err := dock.client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: options.Image,
			Cmd:   options.Cmd,
			Env:   convertEnvMapToSlice(options.Env),
		},
	})
	if err != nil {
		return "", err
	}

	err = dock.client.StartContainer(container.ID, &docker.HostConfig{
		Binds: options.VolumeBinds,
	})
	if err != nil {
		return "", err
	}

	if !options.Detach {
		_, err = dock.client.WaitContainer(container.ID)
		if err != nil {
			return "", err
		}
	}
	return container.ID, nil
}

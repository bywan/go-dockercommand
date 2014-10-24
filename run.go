package dockercommand

import docker "github.com/fsouza/go-dockerclient"

type RunOptions struct {
	Image       string
	Cmd         []string
	VolumeBinds []string
	Detach      bool
	Env         map[string]string
}

func (dock *Docker) Run(options *RunOptions) (error, string) {
	client, err := docker.NewClient(ResolveDockerEndpoint(dock.endpointURL))
	if err != nil {
		return err, ""
	}
	if err = PullImageIfNotExist(options.Image); err != nil {
		return err, ""
	}

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: options.Image,
			Cmd:   options.Cmd,
		},
	})
	if err != nil {
		return err, ""
	}

	err = client.StartContainer(container.ID, &docker.HostConfig{
		Binds: options.VolumeBinds,
	})
	if err != nil {
		return err, ""
	}

	if !options.Detach {
		_, err = client.WaitContainer(container.ID)
		if err != nil {
			return err, ""
		}
	}
	return nil, container.ID
}

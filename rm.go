package dockercommand

import docker "github.com/fsouza/go-dockerclient"

type RmOptions struct {
	Container []string
	Force     bool
}

func (dock *Docker) Rm(options *RmOptions) error {
	client, err := docker.NewClient(ResolveDockerEndpoint(dock.endpointURL))
	if err != nil {
		return err
	}

	for _, containerID := range options.Container {
		err = client.RemoveContainer(docker.RemoveContainerOptions{ID: containerID, Force: options.Force})
		if err != nil {
			return err
		}
	}
	return nil
}

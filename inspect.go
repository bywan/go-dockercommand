package dockercommand

import docker "github.com/fsouza/go-dockerclient"

func (dock *Docker) Inspect(containerID string) (error, *docker.Container) {
	client, err := docker.NewClient(resolveDockerEndpoint(dock.endpointURL))
	if err != nil {
		return err, nil
	}

	container, err := client.InspectContainer(containerID)
	if err != nil {
		return err, nil
	}

	return nil, container
}

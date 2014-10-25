package dockercommand

import docker "github.com/fsouza/go-dockerclient"

func (dock *Docker) Inspect(containerID string) (*docker.Container, error) {
	client, err := docker.NewClient(resolveDockerEndpoint(dock.endpointURL))
	if err != nil {
		return nil, err
	}

	container, err := client.InspectContainer(containerID)
	if err != nil {
		return nil, err
	}

	return container, nil
}

package dockercommand

import docker "github.com/fsouza/go-dockerclient"

type PsOptions struct {
	All    bool
	Size   bool
	Limit  int
	Since  string
	Before string
}

func (dock *Docker) Ps(options *PsOptions) (error, []Container) {
	client, err := docker.NewClient(resolveDockerEndpoint(dock.endpointURL))
	if err != nil {
		return err, nil
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{
		All:    options.All,
		Size:   options.Size,
		Limit:  options.Limit,
		Since:  options.Since,
		Before: options.Before,
	})
	if err != nil {
		return err, nil
	}

	return nil, convertAllContainers(containers)
}

func convertAllContainers(APIContainers []docker.APIContainers) []Container {
	var output []Container
	for _, container := range APIContainers {
		output = append(output, convertContainer(container))
	}
	return output
}

func convertContainer(APIContainer docker.APIContainers) Container {
	return Container{
		ID:         APIContainer.ID,
		Image:      APIContainer.Image,
		Command:    APIContainer.Command,
		Created:    APIContainer.Created,
		Status:     APIContainer.Status,
		Ports:      convertAllPorts(APIContainer.Ports),
		SizeRw:     APIContainer.SizeRw,
		SizeRootFs: APIContainer.SizeRootFs,
		Names:      APIContainer.Names,
	}
}

func convertAllPorts(APIPorts []docker.APIPort) []Port {
	var output []Port
	for _, port := range APIPorts {
		output = append(output, convertPort(port))
	}
	return output
}

func convertPort(APIPort docker.APIPort) Port {
	return Port{
		PrivatePort: APIPort.PrivatePort,
		PublicPort:  APIPort.PublicPort,
		Type:        APIPort.Type,
		IP:          APIPort.IP,
	}
}

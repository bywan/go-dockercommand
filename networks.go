package dockercommand

import docker "github.com/fsouza/go-dockerclient"

func (dock *Docker) Networks() ([]docker.Network, error) {
	return dock.client.ListNetworks()
}

func (dock *Docker) CreateNetwork(opts docker.CreateNetworkOptions) (*docker.Network, error) {
	return dock.client.CreateNetwork(opts)
}

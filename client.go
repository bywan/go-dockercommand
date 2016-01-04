package dockercommand

import docker "github.com/fsouza/go-dockerclient"

type Client interface {
	ListContainers(opts docker.ListContainersOptions) ([]docker.APIContainers, error)
	ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error)
	BuildImage(opts docker.BuildImageOptions) error
	InspectContainer(id string) (*docker.Container, error)
	PullImage(opts docker.PullImageOptions, auth docker.AuthConfiguration) error
	RemoveContainer(opts docker.RemoveContainerOptions) error
	CreateContainer(opts docker.CreateContainerOptions) (*docker.Container, error)
	StartContainer(id string, hostConfig *docker.HostConfig) error
	WaitContainer(id string) (int, error)
	StopContainer(id string, timeout uint) error
	InspectImage(name string) (*docker.Image, error)
	Logs(opts docker.LogsOptions) error
	ListNetworks() ([]docker.Network, error)
	CreateNetwork(docker.CreateNetworkOptions) (*docker.Network, error)
}

type FsouzaClient struct {
	client *docker.Client
}

func (c *FsouzaClient) Logs(opts docker.LogsOptions) error {
	return c.client.Logs(opts)
}

func (c *FsouzaClient) InspectImage(name string) (*docker.Image, error) {
	return c.client.InspectImage(name)
}

func (c *FsouzaClient) StopContainer(id string, timeout uint) error {
	return c.client.StopContainer(id, timeout)
}

func (c *FsouzaClient) WaitContainer(id string) (int, error) {
	return c.client.WaitContainer(id)
}

func (c *FsouzaClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	return c.client.StartContainer(id, hostConfig)
}

func (c *FsouzaClient) CreateContainer(opts docker.CreateContainerOptions) (*docker.Container, error) {
	return c.client.CreateContainer(opts)
}

func (c *FsouzaClient) ListContainers(opts docker.ListContainersOptions) ([]docker.APIContainers, error) {
	return c.client.ListContainers(opts)
}

func (c *FsouzaClient) ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error) {
	return c.client.ListImages(opts)
}

func (c *FsouzaClient) BuildImage(opts docker.BuildImageOptions) error {
	return c.client.BuildImage(opts)
}

func (c *FsouzaClient) InspectContainer(id string) (*docker.Container, error) {
	return c.client.InspectContainer(id)
}

func (c *FsouzaClient) PullImage(opts docker.PullImageOptions, auth docker.AuthConfiguration) error {
	return c.client.PullImage(opts, auth)
}

func (c *FsouzaClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	return c.client.RemoveContainer(opts)
}

func (c *FsouzaClient) ListNetworks() ([]docker.Network, error) {
	return c.client.ListNetworks()
}

func (c *FsouzaClient) CreateNetwork(opts docker.CreateNetworkOptions) (*docker.Network, error) {
	return c.client.CreateNetwork(opts)
}

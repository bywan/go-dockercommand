package mocks

import "github.com/fsouza/go-dockerclient"

type FakeClient struct {
}

func (c *FakeClient) Logs(opts docker.LogsOptions) error {
	panic("Not Implemented Yet")
}

func (c *FakeClient) InspectImage(name string) (*docker.Image, error) {
	panic("Not Implemented Yet")
}

func (c *FakeClient) StopContainer(id string, timeout uint) error {
	panic("Not Implemented Yet")
}

func (c *FakeClient) WaitContainer(id string) (int, error) {
	panic("Not Implemented Yet")
}

func (c *FakeClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	panic("Not Implemented Yet")
}

func (c *FakeClient) CreateContainer(opts docker.CreateContainerOptions) (*docker.Container, error) {
	panic("Not Implemented Yet")
}

func (c *FakeClient) ListContainers(opts docker.ListContainersOptions) ([]docker.APIContainers, error) {
	return []docker.APIContainers{
		{
			ID:    "8dfafdbc3a40",
			Image: "ubuntu:latest",
		}, {
			ID:    "0236fd017853",
			Image: "base:2",
		}}, nil
}

func (c *FakeClient) ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error) {
	return []docker.APIImages{
		{
			ID: "8dfafdbc3a40",
		}, {
			ID: "0236fd017853",
		}}, nil
}

func (c *FakeClient) BuildImage(opts docker.BuildImageOptions) error {
	panic("Not Implemented Yet")
}

func (c *FakeClient) InspectContainer(id string) (*docker.Container, error) {
	panic("Not Implemented Yet")
}

func (c *FakeClient) PullImage(opts docker.PullImageOptions, auth docker.AuthConfiguration) error {
	panic("Not Implemented Yet")
}

func (c *FakeClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	panic("Not Implemented Yet")
}

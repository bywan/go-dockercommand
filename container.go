package dockercommand

import (
	"io"
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

type Container struct {
	info   *docker.Container
	client *docker.Client
}

func (c *Container) Wait() (int, error) {
	return c.client.WaitContainer(c.info.ID)
}

func (c *Container) StreamLogs(w io.Writer) {
	options := docker.LogsOptions{
		Container:    c.info.ID,
		OutputStream: w,
		ErrorStream:  w,
		Follow:       true,
		Stdout:       true,
		Stderr:       true,
	}
	go func() {
		err := c.client.Logs(options)
		if err != nil {
			log.Println(err.Error())
		}
	}()
}

func (c *Container) Remove() error {
	options := docker.RemoveContainerOptions{
		ID:    c.info.ID,
		Force: true,
	}
	return c.client.RemoveContainer(options)
}

func (c *Container) ID() string {
	return c.info.ID
}

func (c *Container) Inspect() (*docker.Container, error) {
	return c.client.InspectContainer(c.info.ID)
}

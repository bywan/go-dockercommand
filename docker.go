package dockercommand

import (
	docker "github.com/fsouza/go-dockerclient"
	"os"
)

type Docker struct {
	client *docker.Client
}

func NewDocker(endpoint string) (*Docker, error) {
	endpoint = resolveDockerEndpoint(endpoint)
	client, err := docker.NewClient(resolveDockerEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return &Docker{client}, nil
}

func resolveDockerEndpoint(input string) string {
	if len(input) != 0 {
		return input
	}
	if len(os.Getenv("DOCKER_HOST")) != 0 {
		return os.Getenv("DOCKER_HOST")
	}
	return "unix:///var/run/docker.sock"
}

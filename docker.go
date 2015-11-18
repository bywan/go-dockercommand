package dockercommand

import (
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

type Docker struct {
	client Client
}

func NewDocker(endpoint string) (*Docker, error) {
	endpoint = resolveDockerEndpoint(endpoint)

	if len(os.Getenv("DOCKER_CERT_PATH")) != 0 {
		client, err := docker.NewTLSClient(endpoint, os.Getenv("DOCKER_CERT_PATH")+"/cert.pem",
			os.Getenv("DOCKER_CERT_PATH")+"/key.pem",
			os.Getenv("DOCKER_CERT_PATH")+"/ca.pem")
		if err != nil {
			log.Fatal(err)
		}
		return &Docker{&FsouzaClient{client}}, nil
	}

	client, err := docker.NewClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &Docker{&FsouzaClient{client}}, nil
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

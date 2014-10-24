package dockercommand

import (
	"fmt"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func ResolveDockerEndpoint(input string) string {
	if len(input) != 0 {
		return input
	}
	if len(os.Getenv("DOCKER_HOST")) != 0 {
		return os.Getenv("DOCKER_HOST")
	}
	return "unix:///var/run/docker.sock"
}

func PullImageIfNotExist(image string) error {
	client, err := docker.NewClient(ResolveDockerEndpoint(""))
	if err != nil {
		return err
	}
	_, err = client.InspectImage(image)
	fmt.Printf("1\n")
	if err != nil && err.Error() == "no such image" {
		fmt.Printf("2\n")
		err = client.PullImage(docker.PullImageOptions{Repository: image, OutputStream: os.Stdout},
			docker.AuthConfiguration{})
	}
	return err
}

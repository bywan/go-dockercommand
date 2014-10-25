package dockercommand

import (
	"fmt"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func resolveDockerEndpoint(input string) string {
	if len(input) != 0 {
		return input
	}
	if len(os.Getenv("DOCKER_HOST")) != 0 {
		return os.Getenv("DOCKER_HOST")
	}
	return "unix:///var/run/docker.sock"
}

func pullImageIfNotExist(image string) error {
	client, err := docker.NewClient(resolveDockerEndpoint(""))
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

func convertEnvMapToSlice(envMap map[string]string) []string {
	envSlice := []string{}
	for key, value := range envMap {
		envSlice = append(envSlice, fmt.Sprintf("%s=%s", key, value))
	}
	return envSlice
}

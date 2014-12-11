package dockercommand

import (
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

type PullOptions struct {
	Image string
}

func (dock *Docker) Pull(options *PullOptions) error {
	return dock.client.PullImage(docker.PullImageOptions{Repository: options.Image, OutputStream: os.Stdout},
		docker.AuthConfiguration{})
}

package dockercommand

import (
	"os"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

type PullOptions struct {
	Image string
}

func (dock *Docker) Pull(options *PullOptions) error {
	var image string
	var tag string
	if strings.Contains(options.Image, ":") {
		imageSplit := strings.Split(options.Image, ":")
		image = imageSplit[0]
		tag = imageSplit[1]
	} else {
		image = options.Image
		tag = "latest"
	}
	return dock.client.PullImage(docker.PullImageOptions{
		Repository:   image,
		Tag:          tag,
		OutputStream: os.Stdout,
	}, docker.AuthConfiguration{})
}

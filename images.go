package dockercommand

import docker "github.com/fsouza/go-dockerclient"

type ImagesOptions struct {
	All     bool
	Filters map[string][]string
	Digests bool
}

func (dock *Docker) Images(options *ImagesOptions) ([]docker.APIImages, error) {
	return dock.client.ListImages(docker.ListImagesOptions{
		All:     options.All,
		Filters: options.Filters,
		Digests: options.Digests,
	})
}

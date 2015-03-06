package dockercommand

import "github.com/bywan/go-dockercommand/mocks"

func NewDockerForTest() (*Docker, error) {
	return &Docker{&mocks.FakeClient{}}, nil
}

func cleanContainer(docker *Docker, ID string) error {
	docker.Stop(&StopOptions{
		ID:      ID,
		Timeout: 1,
	})

	err := docker.Rm(&RmOptions{
		Container:     []string{ID},
		Force:         true,
		RemoveVolumes: true,
	})

	return err
}

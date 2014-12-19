package dockercommand

import "github.com/bywan/go-dockercommand/mocks"

func NewDockerForTest() (*Docker, error) {
	return &Docker{&mocks.FakeClient{}}, nil
}

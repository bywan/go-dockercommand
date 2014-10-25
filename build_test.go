package dockercommand

import "testing"

func TestDockerBuild(t *testing.T) {
	docker := &Docker{}
	err := docker.Build(&BuildOptions{
		Dockerfile: "Dockerfile_test",
		Path:       ".",
		Tag:        "testgodockercommand",
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

}

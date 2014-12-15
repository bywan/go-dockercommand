package dockercommand

import "testing"

func TestDockerPull(t *testing.T) {
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	err = docker.Pull(&PullOptions{
		Image: "mongo:2.4",
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

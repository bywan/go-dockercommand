package dockercommand

import "testing"

func TestDockerPull(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	err = docker.Pull(&PullOptions{
		Image: "ubuntu",
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

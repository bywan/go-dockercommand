package dockercommand

import "testing"

func TestDockerRm(t *testing.T) {
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	container, err := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"ls", "/"},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = container.Remove()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

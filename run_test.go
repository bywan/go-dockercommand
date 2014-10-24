package dockercommand

import "testing"

func TestDockerRun(t *testing.T) {
	docker := &Docker{}
	err, _ := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"echo", "Hello"},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

package dockercommand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerRun(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	container, err := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"/bin/sh", "-c", "while true; do echo hello world; sleep 1; done"},
		Env: map[string]string{
			"TOTO1": "toto1",
			"TOTO2": "toto2",
		},
		VolumeBinds: []string{
			"/:/volumes",
		},
		Detach: true,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = cleanContainer(docker, container.ID())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestDockerRunTwice(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	container1, err := docker.Run(&RunOptions{
		Name:   "test_docker_command",
		Image:  "ubuntu",
		Cmd:    []string{"ls"},
		Detach: true,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	_, err = docker.Run(&RunOptions{
		Name:   "test_docker_command",
		Image:  "ubuntu",
		Cmd:    []string{"ps"},
		Detach: true,
	})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "container already exists")

	err = cleanContainer(docker, container1.ID())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

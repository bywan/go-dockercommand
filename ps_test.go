package dockercommand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortDockerPs(t *testing.T) {
	docker, err := NewDockerForTest()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	containers, err := docker.Ps(&PsOptions{All: true})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	assert.NotEmpty(t, containers)
	assert.Len(t, containers, 2)
	assert.Equal(t, containers[0].ID, "8dfafdbc3a40")
	assert.Equal(t, containers[1].ID, "0236fd017853")
}

func TestLongDockerPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
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

	containers, err := docker.Ps(&PsOptions{
		All: true,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	assert.NotEmpty(t, containers)

	err = cleanContainer(docker, container.ID())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

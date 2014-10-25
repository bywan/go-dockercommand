package dockercommand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerPs(t *testing.T) {
	docker := &Docker{}
	err, _ := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"ls", "/"},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err, containers := docker.Ps(&PsOptions{})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	assert.NotEmpty(t, containers)

}

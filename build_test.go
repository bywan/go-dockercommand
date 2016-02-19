package dockercommand

import "testing"

func TestDockerBuild(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	err = docker.Build(&BuildOptions{
		Dockerfile: "Dockerfile_test",
		Path:       ".",
		Tag:        "testgodockercommand",
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// TODO clean image
}

func TestDockerBuildWithBuildArg(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	err = docker.Build(&BuildOptions{
		Dockerfile: "Dockerfile_test",
		Path:       ".",
		Tag:        "testgodockercommand",
		BuildArgs: map[string]string{
			"HTTP_PROXY": "http://10.20.30.2:1234",
		},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestDockerBuildWithoutCache(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	err = docker.Build(&BuildOptions{
		Dockerfile: "Dockerfile_test",
		Path:       ".",
		NoCache:    true,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
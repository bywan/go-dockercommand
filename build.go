package dockercommand

import (
	"bufio"
	"io"
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

type BuildOptions struct {
	// Relative path of the Dockerfile from the Path
	Dockerfile string
	// The context dir for the docker build
	Path string
	// The repository name
	Tag string
	//The Build args
	BuildArgs map[string]string
}

func (dock *Docker) Build(options *BuildOptions) error {
	return dock.BuildWithLogger(options, log.New(os.Stdout, "", 0))
}

func (dock *Docker) BuildWithLogger(options *BuildOptions, logger *log.Logger) error {
	logsReader, outputbuf := io.Pipe()
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			logger.Printf("%s \n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			logger.Println("There was an error with the scanner in attached container", err)
		}
	}(logsReader)

	buildArgs := make([]docker.BuildArg, len(options.BuildArgs))
	i := 0

	for k, v := range options.BuildArgs {
		buildArgs[i] = docker.BuildArg{
			Name:  k,
			Value: v,
		}
		i++
	}

	opts := docker.BuildImageOptions{
		Name:         options.Tag,
		Dockerfile:   options.Dockerfile,
		ContextDir:   options.Path,
		OutputStream: outputbuf,
		BuildArgs:    buildArgs,
	}
	err := dock.client.BuildImage(opts)

	if err != nil {
		return err
	}
	return nil
}

package dockercommand

import (
	"bufio"
	"fmt"
	"io"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

type RunOptions struct {
	Image       string
	Cmd         []string
	VolumeBinds []string
	Detach      bool
	Logs        bool
	Env         map[string]string
}

func (dock *Docker) Run(options *RunOptions) (string, error) {
	if err := dock.pullImageIfNotExist(options.Image); err != nil {
		return "", err
	}

	container, err := dock.client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: options.Image,
			Cmd:   options.Cmd,
			Env:   convertEnvMapToSlice(options.Env),
		},
	})
	if err != nil {
		return "", err
	}

	err = dock.client.StartContainer(container.ID, &docker.HostConfig{
		Binds: options.VolumeBinds,
	})
	if err != nil {
		return "", err
	}

	if options.Logs {
		r, w := io.Pipe()
		options := docker.LogsOptions{
			Container:    container.ID,
			OutputStream: w,
			ErrorStream:  w,
			Follow:       true,
			Stdout:       true,
			Stderr:       true,
		}
		go func() {
			fmt.Println("Logs container %s", container.ID)
			err := dock.client.Logs(options)
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
		go func(reader io.Reader) {
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				fmt.Printf("%s \n", scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "There was an error with the scanner in attached container", err)
			}
		}(r)
	}

	if !options.Detach {
		_, err = dock.client.WaitContainer(container.ID)
		if err != nil {
			return "", err
		}
	}
	return container.ID, nil
}

go-dockercommand
================

[![Travis](http://img.shields.io/travis/bywan/go-dockercommand.svg?style=flat-square)](https://travis-ci.org/bywan/go-dockercommand)

Go Docker Command is a Go library that provides a fluent interface to manage Docker
containers, in the manner of the Docker CLI

We use the great [fsouza/go-dockerclient](https://github.com/fsouza/go-dockerclient)
to communicate with the Docker engine.

Note: Requires at least go 1.3

# Examples

Start a nginx container with go-dockercommand

```go
package main

import (
	"fmt"

	docker "github.com/bywan/go-dockercommand"
)

func main() {
	// Create go-dockercommand client
	client, err := docker.NewDocker("")
	if err != nil {
		fmt.Printf("Error creating go-dockercommand client, error is: %v", err)
	}

	// Create a nginx container
	_, err = client.Run(&docker.RunOptions{
		Name:   "test",
		Image:  "nginx",
		Detach: true,
	})

	if err != nil {
		fmt.Printf("Error Starting Docker Container, error is: %v", err)
	}
}
```

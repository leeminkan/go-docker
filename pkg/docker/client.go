package docker

import (
	"github.com/docker/docker/client"
)

// Client docker client
var Client *Docker

// Docker docker client
type Docker struct {
	*client.Client
}

// NewDocker create new docker client
func NewDocker() *Docker {

	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	Client = &Docker{client}

	return Client
}

// Setup Initialize the docker client
func Setup() {

	NewDocker()
}

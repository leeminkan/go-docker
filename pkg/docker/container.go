package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListContainers(client *client.Client) ([]types.Container, error) {
	ctx := context.Background()

	// Define the options to use for get image list
	// https://godoc.org/github.com/docker/docker/api/types#Container
	options := types.ContainerListOptions{}

	// Get List Image
	result, err := client.ContainerList(
		ctx,
		options,
	)

	if err != nil {
		return result, err
	}

	return result, err
}

func GetContainer(client *client.Client, containerID string) (types.ContainerJSON, error) {
	ctx := context.Background()

	// Get List Container
	result, err := client.ContainerInspect(
		ctx,
		containerID,
	)

	if err != nil {
		return result, err
	}

	return result, err
}

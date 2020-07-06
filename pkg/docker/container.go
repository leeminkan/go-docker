package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListContainers(client *client.Client) ([]types.Container, error) {
	ctx := context.Background()

	// Define the options to use for get image list
	// https://godoc.org/github.com/docker/docker/api/types#Container
	options := types.ContainerListOptions{
		All: true,
	}

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

func CreateContainerWithName(client *client.Client, containerName string, imageName string) (container.ContainerCreateCreatedBody, error) {
	ctx := context.Background()
	var config container.Config

	config = container.Config{Image: imageName}

	result, err := client.ContainerCreate(ctx, &config, nil, nil, containerName)

	if err != nil {
		return result, err
	}

	return result, err
}

func RemoveContainer(client *client.Client, containerID string) error {
	ctx := context.Background()

	// Define the options to use for get image list
	// https://godoc.org/github.com/docker/docker/api/types#ContainerRemoveOptions
	options := types.ContainerRemoveOptions{
		Force: true,
	}

	err := client.ContainerRemove(
		ctx,
		containerID,
		options,
	)

	return err
}

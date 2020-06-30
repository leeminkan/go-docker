package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListImages(client *client.Client) ([]types.ImageSummary, error) {
	ctx := context.Background()

	// Define the options to use for get image list
	// https://godoc.org/github.com/docker/docker/api/types#ImageListOptions
	options := types.ImageListOptions{}

	// Get List Image
	result, err := client.ImageList(
		ctx,
		options,
	)

	if err != nil {
		return result, err
	}

	return result, err
}

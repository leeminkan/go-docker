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
	imageListOptions := types.ImageListOptions{}

	// Get List Image
	images, err := client.ImageList(
		ctx,
		imageListOptions,
	)

	if err != nil {
		return images, err
	}

	return images, err
}

package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"

	"mime/multipart"

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

func BuildImageFromDockerFile(client *client.Client, tags []string, file multipart.File, fileHeader *multipart.FileHeader) (types.ImageBuildResponse, error) {
	ctx := context.Background()
	var result types.ImageBuildResponse

	// Create a buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Make a TAR header for the file
	tarHeader := &tar.Header{
		Name: fileHeader.Filename,
		Size: int64(fileHeader.Size),
	}

	// Writes the header described for the TAR file
	err := tw.WriteHeader(tarHeader)

	if err != nil {
		return result, err
	}

	// Writes the docker file for the TAR file
	if _, err := io.Copy(tw, file); err != nil {
		return result, err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the options to use for get image list
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	options := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: fileHeader.Filename,
		Remove:     true,
		Tags:       tags,
	}

	// Build the actual image
	result, err = client.ImageBuild(
		ctx,
		dockerFileTarReader,
		options,
	)

	if err != nil {
		return result, err
	}

	// Read the STDOUT from the build process
	defer result.Body.Close()
	_, err = io.Copy(os.Stdout, result.Body)

	if err != nil {
		return result, err
	}

	return result, err
}

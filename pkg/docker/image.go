package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"go-docker/pkg/logging"
	"io"
	"mime/multipart"
	"strings"

	"io/ioutil"

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

func BuildImageFromDockerFile(client *client.Client, tags []string, file multipart.File, fileHeader *multipart.FileHeader) ([]interface{}, error) {
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
		return nil, err
	}

	// Writes the docker file for the TAR file
	if _, err := io.Copy(tw, file); err != nil {
		return nil, err
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
		return nil, err
	}
	defer result.Body.Close()

	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		logging.Warn("Error: %s", err)
	}
	var mResponse = string(response)
	rawData := (strings.Split(mResponse, "\r\n"))
	var mOutput []interface{}
	for _, d := range rawData {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(d), &data)
		if data != nil {
			mOutput = append(mOutput, data)
		}
	}
	return mOutput, err
}

func RemoveImage(client *client.Client, imageID string) ([]types.ImageDeleteResponseItem, error) {
	ctx := context.Background()

	// Define the options to use for get image list
	// https://godoc.org/github.com/docker/docker/api/types#ImageRemoveOptions
	options := types.ImageRemoveOptions{
		Force: true,
	}

	// Get List Image
	result, err := client.ImageRemove(
		ctx,
		imageID,
		options,
	)

	if err != nil {
		return nil, err
	}

	return result, err
}

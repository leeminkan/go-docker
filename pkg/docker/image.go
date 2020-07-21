package docker

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-docker/models"
	"go-docker/pkg/logging"
	imageType "go-docker/type/image"
	"io"
	"mime/multipart"

	"go-docker/service/image_service"

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

func BuildImageFromDockerFile(client *client.Client, mOptions imageType.OptionsBuildImage, file multipart.File, fileHeader *multipart.FileHeader) (types.ImageBuildResponse, error) {
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
		Tags:       mOptions.Tags,
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

	return result, err
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

func GetImage(client *client.Client, imageID string) (types.ImageInspect, error) {
	ctx := context.Background()

	// Get List Image
	result, _, err := client.ImageInspectWithRaw(
		ctx,
		imageID,
	)

	if err != nil {
		return result, err
	}

	return result, err
}

func BuildImageFromTar(client *client.Client, mOptions imageType.OptionsBuildImage, file io.Reader) (types.ImageBuildResponse, error) {
	ctx := context.Background()
	var result types.ImageBuildResponse

	// Define the options to use for build image
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	options := types.ImageBuildOptions{
		Context: file,
		Remove:  true,
		Tags:    mOptions.Tags,
	}

	// Build the actual image
	result, err := client.ImageBuild(
		ctx,
		file,
		options,
	)

	if err != nil {
		return result, err
	}

	return result, err
}

func HandleResultForBuild(result io.ReadCloser, image models.ImageBuild) {
	defer result.Close()
	scanner := bufio.NewScanner(result)
	done := false
	var imageID string
	for scanner.Scan() {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(scanner.Text()), &data)
		if data != nil {
			if val, ok := data["aux"]; ok {
				aux := val.(map[string]interface{})
				if id, ok := aux["ID"]; ok {
					logging.Warn("Value: %s", id.(string))
					imageID = id.(string)
					done = true
				}
			}
			fmt.Println(data)
		}
	}
	if done == true {
		image.Update(image.RepoName, imageID, image.UserID, image_service.Status["Done"])
	} else {
		image.Update(image.RepoName, imageID, image.UserID, image_service.Status["Fail"])
	}
}

func PushImage(client *client.Client, image string, registryAuth string) (interface{}, error) {
	ctx := context.Background()

	// Define the options to use for push image
	// https://godoc.org/github.com/docker/docker/api/types#ImagePushOptions
	options := types.ImagePushOptions{
		RegistryAuth: registryAuth,
	}

	// Push image
	result, err := client.ImagePush(
		ctx,
		image,
		options,
	)

	if err != nil {
		return nil, err
	}

	go ReadResult(result)

	return result, err
}

func ReadResult(result io.ReadCloser) {
	defer result.Close()
	scanner := bufio.NewScanner(result)
	for scanner.Scan() {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(scanner.Text()), &data)
		fmt.Println(data)
	}
}

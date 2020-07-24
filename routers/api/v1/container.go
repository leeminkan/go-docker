package v1

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	natting "github.com/docker/go-connections/nat"
)

func runContainer(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {
	// Define a PORT opening
	newport, err := natting.NewPort("tcp", port)
	if err != nil {
		fmt.Println("Unable to create docker port")
		return err
	}

	// Configured hostConfig:
	// https://godoc.org/github.com/docker/docker/api/types/container#HostConfig
	hostConfig := &container.HostConfig{
		PortBindings: natting.PortMap{
			newport: []natting.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{},
		},
	}

	// Define Network config (why isn't PORT in here...?:
	// https://godoc.org/github.com/docker/docker/api/types/network#NetworkingConfig
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	// Define ports to be exposed (has to be same as hostconfig.portbindings.newport)
	exposedPorts := map[natting.Port]struct{}{
		newport: struct{}{},
	}

	// Configuration
	// https://godoc.org/github.com/docker/docker/api/types/container#Config
	config := &container.Config{
		Image:        imagename,
		Env:          inputEnv,
		ExposedPorts: exposedPorts,
		Hostname:     fmt.Sprintf("%s-hostnameexample", imagename),
	}

	// Creating the actual container. This is "nil,nil,nil" in every example.
	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		networkConfig,
		containername,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	// Run the actual container
	client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cont.ID)

	return nil
}

// @Summary Run example container
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /containers/run-example [get]
func RunContainers(c *gin.Context) {
	appG := app.Gin{C: c}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client")
	}

	imagename := "karthequian/helloworld"
	containername := "test_container"
	portopening := "8080"
	inputEnv := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
	err = runContainer(cli, imagename, containername, portopening, inputEnv)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	data := make(map[string]interface{})

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func buildImage(client *client.Client, tags []string, dockerfile string) error {
	ctx := context.Background()

	// Create a buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Create a filereader
	dockerFileReader, err := os.Open(dockerfile)
	if err != nil {
		return err
	}

	// Read the actual Dockerfile
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		return err
	}

	// Make a TAR header for the file
	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readDockerFile)),
	}

	// Writes the header described for the TAR file
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}

	// Writes the dockerfile data to the TAR file
	_, err = tw.Write(readDockerFile)
	if err != nil {
		return err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	// Build the actual image
	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

// @Summary Create example image
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/create-example [get]
func CreateImages(c *gin.Context) {
	appG := app.Gin{C: c}

	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	// Client, imagename and Dockerfile location
	tags := []string{"this_is_a_imagename"}
	dockerfile := "Dockerfile"
	err = buildImage(client, tags, dockerfile)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	data := make(map[string]interface{})

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary Get list containers
// @Produce  json
// @Tags  Containers
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /containers [get]
func GetContainers(c *gin.Context) {
	appG := app.Gin{C: c}

	containers, err := docker.ListContainers(docker.Client.Client)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_CONTAINER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, containers)
	// text := "xxx"
	// tokenPub := Global.Publish("abc", 0, false, text)
	// tokenPub.Wait()
}

// @Summary Get single container
// @Produce  json
// @Tags  Containers
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /containers/{id} [get]
func GetContainer(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")

	result, err := docker.GetContainer(docker.Client.Client, id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

type CreateContainerInput struct {
	Name      string `form:"name" binding:"required"`
	ImageName string `form:"imageName" binding:"required"`
}

// @Summary Create a container
// @Produce  json
// @Accept  multipart/form-data
// @Tags  Containers
// @Param name formData string true "Name"
// @Param imageName formData string false "ImageName"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /containers [post]
func CreateContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form CreateContainerInput
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	result, err := docker.CreateContainerWithName(docker.Client.Client, form.Name, form.ImageName)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

// @Summary Remove container
// @Produce  json
// @Tags  Containers
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /containers/{id} [delete]
func RemoveContainer(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")

	err := docker.RemoveContainer(docker.Client.Client, id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "Successful!")
}

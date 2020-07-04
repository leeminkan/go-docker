package v1

import (
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get list images
// @Produce  json
// @Tags  Images
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/images [get]
func GetImages(c *gin.Context) {
	appG := app.Gin{C: c}

	images, err := docker.ListImages(docker.Client.Client)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, images)
}

// @Summary Build images from docker file
// @Produce  json
// @Accept  multipart/form-data
// @Tags  Images
// @Param file formData file true "Docker File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/images/build-from-docker-file [post]
func BuildImageFromDockerFile(c *gin.Context) {
	appG := app.Gin{C: c}
	file, fileHeader, err := c.Request.FormFile("file")

	defer file.Close()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if file == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tags := []string{"tag", "tag2"}

	response, err := docker.BuildImageFromDockerFile(docker.Client.Client, tags, file, fileHeader)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, response)
}

// @Summary Remove image
// @Produce  json
// @Tags  Images
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/images/{id} [delete]
func RemoveImage(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")

	result, err := docker.RemoveImage(docker.Client.Client, id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

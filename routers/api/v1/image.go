package v1

import (
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	imageType "go-docker/type/image"
	"net/http"

	"go-docker/pkg/util"

	"github.com/gin-gonic/gin"
)

// @Summary Get single image
// @Produce  json
// @Tags  Images
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/{id} [get]
func GetImage(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")

	result, err := docker.GetImage(docker.Client.Client, id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

// @Summary Get list images
// @Produce  json
// @Tags  Images
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images [get]
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
// @Param options query image.OptionsBuildImage true "Options"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/build-from-docker-file [post]
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

	var options imageType.OptionsBuildImage
	err = c.ShouldBindQuery(&options)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	response, err := docker.BuildImageFromDockerFile(docker.Client.Client, options, file, fileHeader)

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
// @Router /images/{id} [delete]
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

// @Summary Build images from tar
// @Produce  json
// @Accept  multipart/form-data
// @Tags  Images
// @Param file formData file true "Tar"
// @Param options query image.OptionsBuildImage true "Options"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/build-from-tar [post]
func BuildImageFromTar(c *gin.Context) {
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
	_, success := util.Find(fileHeader.Header["Content-Type"], "application/x-tar")

	if success != true {
		logging.Warn("Content-type header not properly set in the request. Expected 'application/x-tar', got %s", fileHeader.Header)
		appG.Response(http.StatusBadRequest, e.ERROR_TYPE_MUST_BE_TAR, nil)
		return
	}

	var options imageType.OptionsBuildImage
	err = c.ShouldBindQuery(&options)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	response, err := docker.BuildImageFromTar(docker.Client.Client, options, file)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, response)
}

// @Summary Push Image
// @Produce  json
// @Accept  application/json
// @Tags  Images
// @Param options query image.OptionsPushImage true "Options"
// @Param body body image.InputPushImage true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/push [post]
func PushImage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form imageType.InputPushImage
	)

	var options imageType.OptionsPushImage
	err := c.ShouldBindQuery(&options)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		logging.Warn(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	result, err := docker.PushImage(docker.Client.Client, form.Image, options)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
	return
}

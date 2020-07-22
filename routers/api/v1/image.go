package v1

import (
	"go-docker/models"
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/service/image_service"
	imageType "go-docker/type/image"
	"net/http"

	"go-docker/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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
// @Security ApiKeyAuth
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

	result, err := docker.BuildImageFromDockerFile(docker.Client.Client, options, file, fileHeader)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	user, _ := c.MustGet("user").(models.User)
	imageService := image_service.ImageBuild{
		RepoName: options.Tags[0],
		UserID:   user.ID,
		Status:   image_service.Status["OnProgress"],
	}
	err = imageService.RemoveRepoNameIfExist()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	image, err := imageService.CreateBuild()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	go docker.HandleResultForBuild(result.Body, image)

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

// @Summary Remove image
// @Produce  json
// @Tags  Images
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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

	result, err := docker.BuildImageFromTar(docker.Client.Client, options, file)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	user, _ := c.MustGet("user").(models.User)
	imageService := image_service.ImageBuild{
		RepoName: options.Tags[0],
		UserID:   user.ID,
		Status:   image_service.Status["OnProgress"],
	}
	err = imageService.RemoveRepoNameIfExist()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	image, err := imageService.CreateBuild()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	go docker.HandleResultForBuild(result.Body, image)

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

// @Summary Push Image
// @Produce  json
// @Accept  application/json
// @Security ApiKeyAuth
// @Tags  Images
// @Param body body image.InputPushImage true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/push [post]
func PushImage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form imageType.InputPushImage
	)

	user, _ := c.MustGet("user").(models.User)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		logging.Warn(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	result, err := docker.PushImage(docker.Client.Client, form.Image, user.XRegistryAuth)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
	return
}

// @Summary Get image build
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Images
// @Param options query image.InputGetImageBuild true "Options"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-build [get]
func GetImageBuild(c *gin.Context) {
	appG := app.Gin{C: c}

	var options imageType.InputGetImageBuild

	err := c.ShouldBindQuery(&options)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	imageService := image_service.ImageBuild{
		RepoName: options.Tag,
		ImageID:  options.Image,
	}

	_, image, err := imageService.Get()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, image)
}

// @Summary Get image build by id
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Images
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-build/{id} [get]
func GetImageBuildByID(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	imageService := image_service.ImageBuild{
		ID: id,
	}
	_, image, err := imageService.GetByID()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, image)
}

// @Summary Tag image
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Images
// @Param body body image.InputTagImage true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images/change-tag [post]
func ChangeTagImage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form imageType.InputTagImage
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	err := docker.TagImage(docker.Client.Client, form.Image, form.Tag)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_DEVICE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

package v1

import (
	"go-docker/models"
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/pkg/util"
	"go-docker/service/image_service"
	imageType "go-docker/type/image"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get image build
// @Produce  json
// @Security ApiKeyAuth
// @Tags  ImageBuild
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
		ImageID:  options.ImageID,
	}

	_, image, err := imageService.Get()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, image)
}

// @Summary Get list image build
// @Produce  json
// @Security ApiKeyAuth
// @Tags  ImageBuild
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-list-build [get]
func GetListImageBuild(c *gin.Context) {
	appG := app.Gin{C: c}

	images, err := image_service.GetListImageBuild()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, images)
}

// @Summary Get image build by id
// @Produce  json
// @Security ApiKeyAuth
// @Tags  ImageBuild
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

// @Summary Build images from docker file
// @Produce  json
// @Accept  multipart/form-data
// @Security ApiKeyAuth
// @Tags  ImageBuild
// @Param file formData file true "Docker File"
// @Param options query image.OptionsBuildImage true "Options"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-build/from-docker-file [post]
func BuildImageFromDockerFile(c *gin.Context) {
	appG := app.Gin{C: c}

	file, fileHeader, err := c.Request.FormFile("file")
	defer file.Close()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if file == nil {
		logging.Warn(file)
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

	appG.Response(http.StatusOK, e.SUCCESS, image)
}

// @Summary Build images from tar
// @Produce  json
// @Accept  multipart/form-data
// @Security ApiKeyAuth
// @Tags  ImageBuild
// @Param file formData file true "Tar"
// @Param options query image.OptionsBuildImage true "Options"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /image-build/from-tar [post]
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

	appG.Response(http.StatusOK, e.SUCCESS, image)
}

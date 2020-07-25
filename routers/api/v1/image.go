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
	check := image_service.CheckExistRepoToRefuse(form.Image)

	if check {
		logging.Warn("Repo Name existed")
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	result, err := docker.PushImage(docker.Client.Client, form.Image, user.XRegistryAuth)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	imageService := image_service.ImagePush{
		RepoName: form.Image,
		UserID:   user.ID,
		Status:   image_service.Status["OnProgress"],
	}

	image, err := imageService.CreatePush()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	go docker.HandleResultForPush(result, image)

	appG.Response(http.StatusOK, e.SUCCESS, image)
	return
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

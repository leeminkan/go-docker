package v1

import (
	"go-docker/models"
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/pkg/util"
	"go-docker/service/image_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get list image push
// @Produce  json
// @Security ApiKeyAuth
// @Tags  ImagePush
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-list-push [get]
func GetListImagePush(c *gin.Context) {
	appG := app.Gin{C: c}

	images, err := image_service.GetListImagePush()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, images)
}

// @Summary Get image push by id
// @Produce  json
// @Security ApiKeyAuth
// @Tags  ImagePush
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-push/{id} [get]
func GetImagePushByID(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	imageService := image_service.ImagePush{
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

// @Summary Push Image From ID
// @Produce  json
// @Accept  application/json
// @Security ApiKeyAuth
// @Tags  ImagePush
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /images-push/from-build-id/{id} [post]
func PushImageFromID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

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

	user, _ := c.MustGet("user").(models.User)

	data, _ := util.DecodeBase64XRegistryAuth(user.XRegistryAuth)

	err = docker.TagImage(docker.Client.Client, image.ImageID, data.Username+"/"+image.RepoName)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	result, err := docker.PushImage(docker.Client.Client, data.Username+"/"+image.RepoName, user.XRegistryAuth)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	checkExists := image_service.CheckExistRepoToRefuse(data.Username + "/" + image.RepoName)

	// if check {
	// 	logging.Warn("Repo Name existed")
	// 	appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	// 	return
	// }

	if checkExists {

		imageUpdating := image_service.ImagePush{
			RepoName: data.Username + "/" + image.RepoName,
			Status:   image_service.Status["OnProgress"],
		}
		imageUpdated, errUpdated := imageUpdating.UpdateStatusPush()

		if errUpdated != nil {
			logging.Warn(errUpdated)
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		go docker.HandleResultForPush(result, imageUpdated)
		appG.Response(http.StatusOK, e.SUCCESS, imageUpdated)
		return
	} else {
		imageServicePush := image_service.ImagePush{
			RepoName: data.Username + "/" + image.RepoName,
			UserID:   user.ID,
			Status:   image_service.Status["OnProgress"],
		}
		imageCreate, errCreate := imageServicePush.CreatePush()

		if errCreate != nil {
			logging.Warn(errCreate)
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}

		go docker.HandleResultForPush(result, imageCreate)

		appG.Response(http.StatusOK, e.SUCCESS, imageCreate)
		return
	}
	// appG.Response(http.StatusOK, e.SUCCESS, result)
	// return
}

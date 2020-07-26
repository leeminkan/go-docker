package v1

import (
	"fmt"
	"go-docker/models"
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/pkg/util"
	"go-docker/service/image_service"
	"go-docker/service/repo_service"
	"go-docker/service/tag_service"
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
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_IMAGE_PUSH, nil)
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
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_IMAGE_PUSH_BY_ID, nil)
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
		appG.Response(http.StatusBadRequest, e.ERROR_IMAGE_BUILD_NOT_FOUND, nil)
		return
	}

	user, _ := c.MustGet("user").(models.User)

	data, _ := util.DecodeBase64XRegistryAuth(user.XRegistryAuth)

	fullRepoName := data.Username + "/" + image.RepoName + ":" + image.Tag

	err = docker.TagImage(docker.Client.Client, image.ImageID, fullRepoName)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_TAG_IMAGE, nil)
		return
	}

	result, err := docker.PushImage(docker.Client.Client, fullRepoName, user.XRegistryAuth)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_PUSH_IMAGE_FAIL, nil)
		return
	}
	var repo models.RepoDockerHub
	var tag models.TagDockerHub
	var imageCreate models.ImagePush

	repoService := repo_service.RepoDockerHub{
		RepoName: data.Username + "/" + image.RepoName,
	}

	exist, repo := repoService.CheckRepoDockerHubExist()

	if exist == true {
		tagService := tag_service.TagDockerHub{
			Tag:    image.Tag,
			RepoID: repo.ID,
		}

		exist, tag = tagService.CheckTagDockerHubExist()

		if exist != true {
			tag, err = tagService.CreateTagDockerHub()

			if err != nil {
				logging.Warn(err)
				appG.Response(http.StatusInternalServerError, e.ERROR_TAG_CREATE, nil)
				return
			}
		} else {
			imageServicePush := image_service.ImagePush{
				TagID:        tag.ID,
				UserID:       user.ID,
				BuildID:      image.ID,
				FullRepoName: fullRepoName,
				Status:       image_service.Status["OnProgress"],
			}

			exist, _ := imageServicePush.CheckImagePushExist()

			if exist == true {
				logging.Warn("Image already push done or on progress")
				appG.Response(http.StatusBadRequest, e.ERROR_PUSH_IMAGE_EXISTED, nil)
				return
			}
		}
	} else {
		repo, err = repoService.CreateRepoDockerHub()

		if err != nil {
			logging.Warn(err)
			appG.Response(http.StatusInternalServerError, e.ERROR_REPO_CREATE, nil)
			return
		}

		tagService := tag_service.TagDockerHub{
			Tag:    image.Tag,
			RepoID: repo.ID,
		}
		tag, err = tagService.CreateTagDockerHub()

		if err != nil {
			logging.Warn(err)
			appG.Response(http.StatusInternalServerError, e.ERROR_TAG_CREATE, nil)
			return
		}
	}

	imageServicePush := image_service.ImagePush{
		TagID:        tag.ID,
		UserID:       user.ID,
		BuildID:      image.ID,
		FullRepoName: fullRepoName,
		Status:       image_service.Status["OnProgress"],
	}
	imageCreate, err = imageServicePush.CreatePush()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	fmt.Println("Van tiep tuc a???")

	go docker.HandleResultForPush(result, imageCreate)

	appG.Response(http.StatusOK, e.SUCCESS, imageCreate)
	return
}

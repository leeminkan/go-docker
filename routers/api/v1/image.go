package v1

import (
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
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

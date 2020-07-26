package v1

import (
	"go-docker/models"
	"go-docker/pkg/app"
	"go-docker/pkg/docker"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"net/http"

	dockerType "go-docker/type/docker"

	b64 "encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// @Summary Docker Hub Login
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Docker
// @Param login body docker.LoginDockerInput true "login"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /docker/login [post]
func LoginDockerHub(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form dockerType.LoginDockerInput
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		logging.Warn(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	result, err := docker.RegistryLogin(docker.Client.Client, form.Username, form.Password)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusUnauthorized, e.ERROR_DOCKER_LOGIN_FAIL, nil)
		return
	}

	user, _ := c.MustGet("user").(models.User)

	mJson, err := json.Marshal(form)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	sEnc := b64.StdEncoding.EncodeToString([]byte(mJson))

	err = user.UpdateXRegistryAuth(true, sEnc)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
	return
}

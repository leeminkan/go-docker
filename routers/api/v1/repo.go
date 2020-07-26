package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"go-docker/pkg/app"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/service/repo_service"
)

// @Summary Get List Repo
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Repos
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /repos [get]
func GetListRepo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	repos, err := repo_service.GetListRepo()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_REPO, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, repos)
	return
}

// @Summary Get Repo By ID
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Repos
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /repos/{id} [get]
func GetRepoByID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()

	repoService := repo_service.RepoDockerHub{
		ID: id,
	}

	exist, repo := repoService.GetRepoByID()

	if exist != true {
		appG.Response(http.StatusBadRequest, e.ERROR_GET_LIST_REPO, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, repo)
	return
}

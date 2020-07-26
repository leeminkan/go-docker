package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"go-docker/pkg/app"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/service/repo_service"
	"go-docker/service/tag_service"
)

// @Summary Get List Tag By Repo ID
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Tags
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags/repo/{id} [get]
func GetListTagByRepoID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()

	repoService := repo_service.RepoDockerHub{
		ID: id,
	}

	exist, _ := repoService.CheckRepoDockerHubExistByID()

	if exist != true {
		appG.Response(http.StatusBadRequest, e.ERROR_REPO_NOT_FOUND, nil)
		return
	}

	tags, err := tag_service.GetListTagByRepoID(id)

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_TAG_BY_REPO_ID, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, tags)
	return
}

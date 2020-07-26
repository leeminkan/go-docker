package tag_service

import (
	"go-docker/models"
)

type TagDockerHub struct {
	ID int

	Tag    string
	RepoID int
}

func (tag *TagDockerHub) CreateTagDockerHub() (models.TagDockerHub, error) {
	return models.CreateTagDockerHub(tag.Tag, tag.RepoID)
}

func (tag *TagDockerHub) CheckTagDockerHubExist() (bool, models.TagDockerHub) {
	return models.CheckTagDockerHubExist(tag.Tag, tag.RepoID)
}

func GetListTagByRepoID(repo_id int) ([]models.TagDockerHub, error) {
	return models.GetListTagDockerHubByRepoID(repo_id)
}

package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type TagDockerHub struct {
	Model

	Tag           string        `json:"tag"`
	RepoID        int           `json:"repo_id"`
	RepoDockerHub RepoDockerHub `gorm:"foreignkey:RepoID"`
}

func CreateTagDockerHub(tag string, repo_id int) (TagDockerHub, error) {
	repoTag := TagDockerHub{
		Tag:    tag,
		RepoID: repo_id,
	}

	if err := db.Create(&repoTag).Error; err != nil {
		logging.Warn(err)
		return repoTag, err
	}
	db.Preload("RepoDockerHub").Find(&repoTag)

	return repoTag, nil
}

func GetListTagDockerHub() ([]TagDockerHub, error) {
	var repos []TagDockerHub
	err := db.Where("deleted_on = ?", 0).Preload("RepoDockerHub").Find(&repos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return repos, err
	}

	return repos, nil
}

func CheckTagDockerHubExist(tag string, repo_id int) (bool, TagDockerHub) {
	var repoTag TagDockerHub
	err := db.Where("tag = ? AND repo_id = ? AND deleted_on = ? ", tag, repo_id, 0).First(&repoTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, repoTag
	}

	if repoTag.ID > 0 {
		return true, repoTag
	}

	return false, repoTag
}

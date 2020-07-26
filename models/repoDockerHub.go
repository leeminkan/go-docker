package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type RepoDockerHub struct {
	Model

	RepoName string `json:"repo_name"`
}

func CreateRepoDockerHub(repo_name string) (RepoDockerHub, error) {
	repo := RepoDockerHub{
		RepoName: repo_name,
	}

	if err := db.Create(&repo).Error; err != nil {
		logging.Warn(err)
		return repo, err
	}

	return repo, nil
}

func GetListRepoDockerHub() ([]RepoDockerHub, error) {
	var repos []RepoDockerHub
	err := db.Where("deleted_on = ?", 0).Find(&repos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return repos, err
	}

	return repos, nil
}

func CheckRepoDockerHubExist(repo_name string) (bool, RepoDockerHub) {
	var repo RepoDockerHub
	err := db.Where("repo_name = ? AND deleted_on = ? ", repo_name, 0).First(&repo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, repo
	}

	if repo.ID > 0 {
		return true, repo
	}

	return false, repo
}

package repo_service

import (
	"go-docker/models"
)

type RepoDockerHub struct {
	ID int

	RepoName string
}

func (repo *RepoDockerHub) CreateRepoDockerHub() (models.RepoDockerHub, error) {
	return models.CreateRepoDockerHub(repo.RepoName)
}

func (repo *RepoDockerHub) CheckRepoDockerHubExist() (bool, models.RepoDockerHub) {
	return models.CheckRepoDockerHubExist(repo.RepoName)
}

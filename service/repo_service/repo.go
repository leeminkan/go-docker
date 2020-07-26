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

func (repo *RepoDockerHub) CheckRepoDockerHubExistByID() (bool, models.RepoDockerHub) {
	return models.CheckRepoDockerHubExistByID(repo.ID)
}

func GetListRepo() ([]models.RepoDockerHub, error) {
	return models.GetListRepoDockerHub()
}

func (repo *RepoDockerHub) GetRepoByID() (bool, models.RepoDockerHub) {
	return models.GetRepoDockerHubByID(repo.ID)
}

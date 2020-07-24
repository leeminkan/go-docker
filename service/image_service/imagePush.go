package image_service

import (
	"go-docker/models"
)

type ImagePush struct {
	ID int

	RepoName string
	UserID   int
	Status   string
}

func (image *ImagePush) CreatePush() (models.ImagePush, error) {
	return models.CreateImagePush(image.RepoName, image.UserID, image.Status)
}

func (image *ImagePush) UpdateStatusPush() (models.ImagePush, error) {
	return models.UpdateStatusImagePush(image.RepoName, image.Status)
}

func GetListImagePush() ([]models.ImagePush, error) {
	return models.GetListImagePush()
}

func CheckExistRepoToRefuse(repo_name string) bool {
	return models.CheckExistRepoToRefuse(repo_name)
}

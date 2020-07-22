package image_service

import (
	"go-docker/models"
)

var Status = map[string]string{
	"OnProgress": "on progress",
	"Done":       "done",
	"Fail":       "fail",
}

type ImageBuild struct {
	ID int

	RepoName string
	ImageID  string
	UserID   int
	Status   string
}

func (image *ImageBuild) CreateBuild() (models.ImageBuild, error) {
	return models.CreateImageBuild(image.RepoName, image.ImageID, image.UserID, image.Status)
}

func (image *ImageBuild) RemoveRepoNameIfExist() error {
	return models.RemoveRepoNameIfExist(image.RepoName)
}

func (image *ImageBuild) GetByID() (bool, models.ImageBuild, error) {
	return models.GetImageBuildByID(image.ID)
}

func (image *ImageBuild) Get() (bool, models.ImageBuild, error) {
	return models.GetImageBuild(image.RepoName, image.ImageID)
}

func GetList() ([]models.ImageBuild, error) {
	return models.GetListImageBuild()
}

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

	RepoName    string
	Tag         string
	ImageID     string
	UserID      int
	Status      string
	OldRepoName string
}

func (image *ImageBuild) CreateBuild() (models.ImageBuild, error) {
	return models.CreateImageBuild(image.RepoName, image.Tag, image.ImageID, image.UserID, image.Status, image.OldRepoName)
}

func (image *ImageBuild) RemoveRepoNameAndTagIfExist() error {
	return models.RemoveRepoNameAndTagIfExist(image.RepoName, image.Tag)
}

func (image *ImageBuild) GetByID() (bool, models.ImageBuild, error) {
	return models.GetImageBuildByID(image.ID)
}

func (image *ImageBuild) Get() (bool, models.ImageBuild, error) {
	return models.GetImageBuild(image.RepoName, image.ImageID)
}

func GetListImageBuild() ([]models.ImageBuild, error) {
	return models.GetListImageBuild()
}

func GetListImageBuildCustom(user_id int) ([]models.ImageBuildCustom, error) {
	return models.GetListImageBuildCustom(user_id)
}

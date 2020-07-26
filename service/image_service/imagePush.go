package image_service

import (
	"go-docker/models"
)

type ImagePush struct {
	ID int

	TagID        int
	UserID       int
	BuildID      int
	FullRepoName string
	Digest       string
	Status       string
}

func (image *ImagePush) CreatePush() (models.ImagePush, error) {
	return models.CreateImagePush(image.TagID, image.UserID, image.BuildID, image.FullRepoName, image.Digest, image.Status)
}

func GetListImagePush() ([]models.ImagePush, error) {
	return models.GetListImagePush()
}

func (image *ImagePush) GetByID() (bool, models.ImagePush, error) {
	return models.GetImagePushByID(image.ID)
}

func (image *ImagePush) CheckImagePushExist() (bool, models.ImagePush) {
	return models.CheckImagePushExist(image.TagID, image.BuildID, Status["Done"], Status["OnProgress"])
}

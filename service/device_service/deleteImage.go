package device_service

import "go-docker/models"

type DeleteImage struct {
	ImageID  int
	DeleteOn int
}

func CheckDeleteImage(imageID int) (bool, error) {
	return models.IsDelete(imageID)
}

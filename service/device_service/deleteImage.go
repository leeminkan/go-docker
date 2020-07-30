package device_service

import "go-docker/models"

type DeleteImage struct {
	ImageID int
}

func CheckDeleteImage(imageID int) (bool, error) {
	return models.IsDeleteImage(imageID)
}

func (t *DeleteImage) GetMachineIDByImageID(string, error) {
	return models.IsDeleteImage(t.ImageID)
}

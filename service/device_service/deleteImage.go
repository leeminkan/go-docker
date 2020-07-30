package device_service

import "go-docker/models"

type DeleteImage struct {
	ImageID  int
	DeleteOn int
}

func CheckDeleteImage(imageID int) (bool, error) {
	return models.IsDeleteImage(imageID)
}

func (t *DeleteImage) GetMachineIDByImageID() (string, error) {
	return models.GetMachineIDByImageID(t.ImageID)
}

func GetImage(imageID int) (models.DeviceImage, error) {
	return models.GetImageByID(imageID)
}

func (t *DeleteImage) UpdateDeleteImage() (models.DeviceImage, error) {
	return models.UpdateDeleteImage(t.ImageID, t.DeleteOn)
}

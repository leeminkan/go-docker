package device_service

import (
	"go-docker/models"
)

type StatusPull struct {
	ImagePullID int
}

func (t *StatusPull) GetImagePull() (models.DeviceImage, error) {
	return models.GetImage(t.ImagePullID)
}

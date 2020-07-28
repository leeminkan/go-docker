package device_service

import (
	"go-docker/models"
)

type Update struct {
	FullRepoName string
	MachineID    string
	ImageID      string
	Status       string
}

func (t *Update) UpdateImagePullStatus() (models.DeviceImage, error) {
	return models.UpdatePull(t.FullRepoName, t.MachineID, t.ImageID, t.Status)
}

package device_service

import (
	"go-docker/models"
)

type ControlRun struct {
	ImagePullID   int
	ContainerName string
}

func (t *ControlRun) CheckValueRun() (models.DeviceImage, bool, error) {
	return models.CheckValueRun(t.ImagePullID, t.ContainerName)
}

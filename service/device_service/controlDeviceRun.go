package device_service

import (
	"go-docker/models"
	"strconv"
)

type ControlRun struct {
	ImagePullID   int
	ContainerName string
}

func (t *ControlRun) CheckValueRun() (models.DeviceImage, bool, error) {
	return models.CheckValueRun(t.ImagePullID, t.ContainerName)
}

func (t *ControlRun) CreateDeviceRun() (models.DeviceContainer, error) {
	return models.CreateRun(strconv.Itoa(t.ImagePullID), t.ContainerName)
}

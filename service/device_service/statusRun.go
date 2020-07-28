package device_service

import (
	"go-docker/models"
)

type StatusRun struct {
	ContainerID int
}

func (t *StatusRun) GetContainerRun() (models.DeviceContainer, error) {
	return models.GetContainer(t.ContainerID)
}

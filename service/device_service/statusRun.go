package device_service

import (
	"go-docker/models"
)

type StatusRun struct {
	ContainerID int
	Active      string
}

func (t *StatusRun) GetContainerRun() (models.DeviceContainer, error) {
	return models.GetContainer(t.ContainerID)
}

func (t *StatusRun) UpdateStatusContainer() (models.DeviceContainer, error) {
	return models.UpdateStatusContainer(t.ContainerID, t.Active)
}

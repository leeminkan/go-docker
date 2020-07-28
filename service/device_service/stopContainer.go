package device_service

import (
	"go-docker/models"
)

type StopContainer struct {
	ContainerID int
	DeviceID    int
}

func (t *StopContainer) StopContainerByID() error {
	return models.StopContainer(t.ContainerID)
}

func (t *StopContainer) GetMachineIDByContainerID() (string, error) {
	return models.GetMachineIDByContainerID(t.ContainerID)
}

func (t *StopContainer) StopAllContainerByID() error {
	return models.StopContainer(t.DeviceID)
}

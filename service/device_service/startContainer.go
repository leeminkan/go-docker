package device_service

import (
	"go-docker/models"
)

type StartContainer struct {
	ContainerID int
}

func (t *StartContainer) StartContainerByID() error {
	return models.StartContainer(t.ContainerID)
}

func GetContainerStart(containerID int) (models.DeviceContainer, error) {
	return models.GetContainerStop(containerID)
}

func CheckStatusStartStop(containerID int) (models.DeviceContainer, error) {
	return models.CheckStatusStartStop(containerID)
}

func (t *StartContainer) GetMachineIDStop() (string, error) {
	return models.GetMachineIDByContainerID(t.ContainerID)
}

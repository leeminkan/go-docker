package device_service

import "go-docker/models"

type DeleteContainer struct {
	ContainerID int
	DeleteOn    int
}

func (t *DeleteContainer) UpdateDeleteContainer() (models.DeviceContainer, error) {
	return models.UpdateDelete(t.ContainerID, t.DeleteOn)
}

func CheckDeleteContainer(containerID int) (bool, error) {
	return models.IsDelete(containerID)
}

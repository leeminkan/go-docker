package device_service

import (
	"go-docker/models"
)

type UpdateRun struct {
	ContainerName string
	ImagePullID   string
	Status        string
	Active        string
}

func (t *UpdateRun) UpdateContainerRunStatus() (models.DeviceContainer, error) {
	return models.UpdateRun(t.ContainerName, t.ImagePullID, t.Status, t.Active)
}

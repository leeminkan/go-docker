package device_service

import (
	"go-docker/models"
)

type UpdateRun struct {
	ContainerName string
	ImagePullID   int
	Status        string
	Active        string
}

func (t *UpdateRun) UpdateContainerRunStatus() error {
	return models.UpdateRun(t.ContainerName, t.ImagePullID, t.Status, t.Active)
}

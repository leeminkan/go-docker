package device_service

import (
	"go-docker/models"
)

type Update struct {
	FullRepoName string
	MachineID    string
	Status       string
}

func (t *Update) UpdateImagePullStatus() error {
	return models.UpdatePull(t.FullRepoName, t.MachineID, t.Status)
}

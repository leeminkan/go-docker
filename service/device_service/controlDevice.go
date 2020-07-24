package device_service

import (
	"go-docker/models"
)

type Control struct {
	ID         int
	DeviceName string
	OS         string
	MachineID  string
	RepoID     int
}

func (t *Control) ExistDevice() (bool, error) {
	return models.ExistDevice(t.MachineID)
}

func (t *Control) GetRepoNameFromID() (string, error) {
	return models.GetRepoName(t.RepoID)
}

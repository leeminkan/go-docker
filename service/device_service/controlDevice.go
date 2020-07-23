package device_service

import (
	"go-docker/models"
)

type Control struct {
	ID         int
	DeviceName string
	OS         string
	MachineID  string
	RepoName   string
}

func (t *Control) ExistDevice() (bool, error) {
	return models.ExistDevice(t.ID, t.MachineID)
}

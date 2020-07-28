package device_service

import (
	"go-docker/models"
)

type Control struct {
	DeviceID int
	RepoID   int
}

func (t *Control) CheckDevice() (string, bool, error) {
	return models.CheckDevice(t.DeviceID)
}

func (t *Control) GetFullRepoNameFromID() (string, error) {
	return models.GetFullRepoName(t.RepoID)
}

func (t *Control) CreateDevicePull() (models.DeviceImage, error) {
	return models.CreatePull(t.DeviceID, t.RepoID)
}

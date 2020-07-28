package mqtt

import (
	"encoding/json"
)

type ComeinandPull struct {
	MachineID string
	RepoName  string
}

type ComeinandRun struct {
	ContainerName string
	MachineID     string
	RepoName      string
	ImageID       int
}

type ComeinandStopContainer struct {
	MachineID   string
	ContainerID string
}

func SetValueComeinandPull(machineID string, repoName string) ([]byte, error) {
	comeinandRaw := ComeinandPull{
		MachineID: machineID,
		RepoName:  repoName,
	}
	comeinandJSON, err := json.Marshal(comeinandRaw)
	if err != nil {
		return nil, err
	}
	return comeinandJSON, nil
}

func SetValueComeinandRun(containerName string, machineID string, repoName string, imageID int) ([]byte, error) {
	comeinandRaw := ComeinandRun{
		ContainerName: containerName,
		MachineID:     machineID,
		RepoName:      repoName,
		ImageID:       imageID,
	}
	comeinandJSON, err := json.Marshal(comeinandRaw)
	if err != nil {
		return nil, err
	}
	return comeinandJSON, nil
}

func SetValueComeinandStopContainer(machineID string, containerID string) ([]byte, error) {
	comeinandRaw := ComeinandStopContainer{
		MachineID:   machineID,
		ContainerID: containerID,
	}
	comeinandJSON, err := json.Marshal(comeinandRaw)
	if err != nil {
		return nil, err
	}
	return comeinandJSON, nil
}

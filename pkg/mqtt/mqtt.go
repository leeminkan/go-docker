package mqtt

import (
	"encoding/json"
)

type Comeinand struct {
	MachineID string
	RepoName  string
}

func SetValueComeinand(machineID string, repoName string) ([]byte, error) {
	comeinandRaw := Comeinand{
		MachineID: machineID,
		RepoName:  repoName,
	}
	comeinandJSON, err := json.Marshal(comeinandRaw)
	if err != nil {
		return nil, err
	}
	return comeinandJSON, nil
}

package util

import (
	"go-docker/pkg/logging"

	b64 "encoding/base64"
	"encoding/json"
	dockerType "go-docker/type/docker"
)

func DecodeBase64XRegistryAuth(xRegistryAuth string) (*dockerType.LoginDockerInput, error) {

	sDec, _ := b64.StdEncoding.DecodeString(xRegistryAuth)

	data := &dockerType.LoginDockerInput{}
	err := json.Unmarshal(sDec, data)

	if err != nil {
		logging.Warn(err)
		return data, err
	}

	return data, nil
}

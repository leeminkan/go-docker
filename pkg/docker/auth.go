package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

func RegistryLogin(client *client.Client, username string, password string) (registry.AuthenticateOKBody, error) {
	ctx := context.Background()

	config := types.AuthConfig{
		Username: username,
		Password: password,
	}

	// Get List Image
	result, err := client.RegistryLogin(
		ctx,
		config,
	)

	if err != nil {
		return result, err
	}

	return result, err
}

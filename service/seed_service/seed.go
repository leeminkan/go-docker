package seed_service

import (
	"go-docker/models"
)

type Seed struct {
	ID   int
	Name string
}

func (s *Seed) Create() error {
	return models.CreateSeed(s.Name)
}

func (s *Seed) IsSeedExistedByName() (bool, error) {
	return models.IsSeedExistedByName(s.Name)
}

package migration_service

import (
	"go-docker/models"
)

type Migration struct {
	ID   int
	Name string
}

func (m *Migration) Create() error {
	return models.CreateMigration(m.Name)
}

func (m *Migration) IsMigrationExistedByName() (bool, error) {
	return models.IsMigrationExistedByName(m.Name)
}

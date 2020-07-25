package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type Migration struct {
	Model

	Name string `json:"name"`
}

func CreateMigration(name string) error {
	migration := Migration{
		Name: name,
	}
	if err := db.Create(&migration).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func IsMigrationExistedByName(name string) (bool, error) {
	var migration Migration
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&migration).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, err
	}

	if migration.ID > 0 {
		return true, nil
	}

	return false, nil
}

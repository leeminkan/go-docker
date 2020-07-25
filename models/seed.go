package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type Seed struct {
	Model

	Name string `json:"name"`
}

func CreateSeed(name string) error {
	seed := Seed{
		Name: name,
	}
	if err := db.Create(&seed).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func IsSeedExistedByName(name string) (bool, error) {
	var seed Seed
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&seed).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, err
	}

	if seed.ID > 0 {
		return true, nil
	}

	return false, nil
}

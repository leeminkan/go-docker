package models

import (
	"github.com/jinzhu/gorm"
)

func ExistDevice(machineID string) (bool, error) {
	var device Device
	err := db.Select("id").Where("deleted_on = ? AND machine_id = ? ", 0, machineID).First(&device).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if device.ID > 0 {
		return true, nil
	}

	return false, nil
}

package models

import (
	"github.com/jinzhu/gorm"
)

type Control struct {
	Model

	DeviceName string `json:"device_name"`
	OS         string `json:"os"`
	MachineID  string `json:"machine_id"`
	RepoName   string `json:"repo_name"`
}

func ExistDevice(id int, machineID string) (bool, error) {
	var control Control
	err := db.Select("id").Where("id = ? AND deleted_on = ? AND machine_id = ? ", id, 0, machineID).First(&control).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if control.ID > 0 {
		return true, nil
	}

	return false, nil
}

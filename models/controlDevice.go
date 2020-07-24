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

func GetRepoName(repoID int) (string, error) {
	var imagePush ImagePush
	err := db.Select("repo_name").Where("deleted_on = ? AND id = ? ", 0, repoID).First(&imagePush).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	if imagePush.RepoName != "" {
		return imagePush.RepoName, nil
	}

	return "", nil
}

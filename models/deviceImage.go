package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type DeviceImage struct {
	Model

	FullRepoName string `json:"full_repo_name"`
	DeviceID     int    `json:"device_id"`
	Status       string `json:"status" gorm:"type:enum('on progress', 'done', 'fail');default:'on progress'"`
}

func UpdatePull(repoName string, machineID string, status string) error {
	var deviceImage DeviceImage

	type SelectData struct {
		DeviceID     int
		FullRepoName string
	}
	var results []SelectData
	err := db.Table("device").Select("device_image.device_id, device_image.full_repo_name").Joins("left join device_image on device.id = device_image.device_id").Where("device_image.full_repo_name = ? AND device.machine_id = ? AND device.deleted_on = ? AND device_image.deleted_on = ?", repoName, machineID, 0, 0).Scan(&results).Error
	if err != nil || len(results) == 0 {
		var device Device
		db.Where("machine_id = ? AND deleted_on = ?", machineID, 0).First(&device)
		deviceImageCreate := DeviceImage{
			FullRepoName: repoName,
			DeviceID:     device.ID,
			Status:       status,
		}
		errCreate := db.Create(&deviceImageCreate).Error
		if errCreate != nil {
			logging.Warn(errCreate)
			return errCreate
		}
		return nil
	}
	errUpdate := db.Model(&deviceImage).Where("full_repo_name = ? AND device_id = ? AND deleted_on = ? ", repoName, results[0].DeviceID, 0).Updates(DeviceImage{FullRepoName: repoName, DeviceID: results[0].DeviceID, Status: status}).Error

	if errUpdate != nil {
		logging.Warn(errUpdate)
		return errUpdate
	}

	return nil
}

func GetListImages(machineID string) ([]DeviceImage, error) {
	var listImages []DeviceImage
	var device Device
	db.Where("machine_id = ? AND deleted_on = ?", machineID, 0).First(&device)
	err := db.Where("device_id = ? AND deleted_on = ? AND status = ?", device.ID, 0, "done").Find(&listImages).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		logging.Warn(err)
		return listImages, err
	}
	return listImages, err
}

func GetImage(id int) (DeviceImage, error) {
	var deviceImage DeviceImage
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&deviceImage).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		logging.Warn(err)
		return deviceImage, err
	}
	return deviceImage, nil
}

func GetContainer(id int) (DeviceContainer, error) {
	var deviceContainer DeviceContainer
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&deviceContainer).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		logging.Warn(err)
		return deviceContainer, err
	}
	return deviceContainer, nil
}

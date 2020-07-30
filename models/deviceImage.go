package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type DeviceImage struct {
	Model

	FullRepoName string `json:"full_repo_name"`
	DeviceID     int    `json:"device_id"`
	ImageID      string `json:"image_id"`
	Status       string `json:"status" gorm:"type:enum('on progress', 'done', 'fail');default:'on progress'"`
}

func UpdatePull(repoName string, machineID string, imageID string, status string) (DeviceImage, error) {
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
			ImageID:      imageID,
			Status:       status,
		}
		errCreate := db.Create(&deviceImageCreate).Error
		db.Where("full_repo_name = ? AND device_id = ? AND deleted_on = ? ", repoName, device.ID, 0).First(&deviceImage)

		if errCreate != nil {
			logging.Warn(errCreate)
			return deviceImage, errCreate
		}
		return deviceImage, nil
	}

	errUpdate := db.Model(&deviceImage).Where("full_repo_name = ? AND device_id = ? AND deleted_on = ? ", repoName, results[0].DeviceID, 0).Updates(DeviceImage{FullRepoName: repoName, DeviceID: results[0].DeviceID, ImageID: imageID, Status: status}).Error
	db.Where("full_repo_name = ? AND device_id = ? AND deleted_on = ? ", repoName, results[0].DeviceID, 0).First(&deviceImage)
	if errUpdate != nil {
		logging.Warn(errUpdate)
		return deviceImage, errUpdate
	}

	return deviceImage, nil
}

func GetListImages(id int) ([]DeviceImage, error) {
	var listImages []DeviceImage
	var device Device
	db.Where("id = ? AND deleted_on = ?", id, 0).First(&device)
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

func CreatePull(deviceID int, repoID int) (DeviceImage, error) {
	var deviceImage DeviceImage
	var imagePush ImagePush
	errImage := db.Select("full_repo_name").Where("deleted_on = ? AND id = ? ", 0, repoID).First(&imagePush).Error
	if errImage != nil {
		logging.Warn(errImage)
		return deviceImage, errImage
	}
	var device Device
	errDevice := db.Where("deleted_on = ? AND id = ? ", 0, deviceID).First(&device).Error
	if errDevice != nil {
		logging.Warn(errDevice)
		return deviceImage, errDevice
	}

	deviceImageResult, err := UpdatePull(imagePush.FullRepoName, device.MachineID, "", "on progress")
	if err != nil {
		logging.Warn(err)
		return deviceImageResult, err
	}
	return deviceImageResult, nil
}

func IsDeleteImage(id int) (bool, error) {
	type SelectData struct {
		ImageID int
		ID      int
	}
	var results []SelectData
	err := db.Table("device_image").Select("device_container.image_id, device_container.id").Joins("left join device_container on device_image.id = device_container.image_id").Where("device_image.id = ? AND device_image.deleted_on = ? AND device_container.deleted_on = ?", id, 0, 0).Scan(&results).Error

	if err != nil {
		logging.Warn(err)
		return false, err
	}

	if len(results) == 0 {
		return true, nil
	}
	return false, nil
}

func GetMachineIDByImageID(id int) (string, error) {
	var deviceImage DeviceImage
	db.Where("id = ? AND deleted_on = ?", id, 0).First(&deviceImage)
	var device Device
	err := db.Where("id = ? AND deleted_on = ?", deviceImage.DeviceID, 0).First(&device).Error
	if err != nil {
		logging.Warn(err)
		return "", err
	}
	return device.MachineID, nil
}

func GetImageByID(id int) (DeviceImage, error) {
	var deviceImage DeviceImage
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&deviceImage).Error
	if err != nil {
		logging.Warn(err)
		return deviceImage, err
	}
	return deviceImage, nil
}

func UpdateDeleteImage(id int, delete int) (DeviceImage, error) {
	var deviceImage DeviceImage
	err := db.Model(&deviceImage).Where("id = ? AND deleted_on = ? ", id, 0).Update(
		"deleted_on", delete,
	).Error
	if err != nil {
		logging.Warn(err)
		return deviceImage, err
	}
	return deviceImage, nil
}

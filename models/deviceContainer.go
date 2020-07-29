package models

import (
	"errors"
	"go-docker/pkg/logging"
	"strconv"

	"github.com/jinzhu/gorm"
)

type DeviceContainer struct {
	Model

	ContainerName string `json:"container_name"`
	ImageID       int    `json:"image_id"`
	Status        string `json:"status" gorm:"type:enum('on progress', 'done', 'fail');default:'on progress'"`
	Active        string `json:"active" gorm:"type:enum('start', 'starting', 'stop', 'stopping');default:'stop'"`
}

func CheckValueRun(imagePullID int, containerName string) (DeviceImage, bool, error) {
	var deviceImage DeviceImage
	errImage := db.Where("deleted_on = ? AND id = ? AND status = ? ", 0, imagePullID, "done").First(&deviceImage).Error
	if errImage != nil && errImage == gorm.ErrRecordNotFound {
		logging.Warn(errImage)
		return DeviceImage{}, false, errImage
	}

	return deviceImage, true, nil
}

func CreateRun(imagePullID string, containerName string) (DeviceContainer, error) {
	deviceContainer, err := UpdateRun(containerName, imagePullID, "on progress", "starting")
	if err != nil {
		logging.Warn(err)
		return deviceContainer, err
	}
	return deviceContainer, nil
}

func UpdateRun(containerName string, imagePullID string, status string, active string) (DeviceContainer, error) {
	var deviceContainer DeviceContainer
	err := db.Where("image_id = ? AND container_name = ? AND deleted_on = ? ", imagePullID, containerName, 0).First(&deviceContainer).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		cvtInt, _ := strconv.Atoi(imagePullID)
		deviceContainerCreate := DeviceContainer{
			ContainerName: containerName,
			ImageID:       cvtInt,
			Status:        status,
			Active:        active,
		}
		errCreate := db.Create(&deviceContainerCreate).Error
		db.Where("image_id = ? AND container_name = ? AND deleted_on = ? ", imagePullID, containerName, 0).First(&deviceContainer)

		if errCreate != nil {
			logging.Warn(errCreate)
			return deviceContainer, errCreate
		}
		return deviceContainer, nil
	}
	errUpdate := db.Model(&deviceContainer).Where("image_id = ? AND container_name = ? AND deleted_on = ? ", imagePullID, containerName, 0).Updates(
		DeviceContainer{
			Status: status,
			Active: active,
		}).Error
	db.Where("image_id = ? AND container_name = ? AND deleted_on = ? ", imagePullID, containerName, 0).First(&deviceContainer)

	if errUpdate != nil {
		logging.Warn(errUpdate)
		return deviceContainer, errUpdate
	}
	return deviceContainer, nil
}

func GetListContainers(id int) ([]DeviceContainer, error) {
	var listContainers []DeviceContainer
	var device Device
	db.Where("id = ? AND deleted_on = ?", id, 0).First(&device)
	err := db.Table("device_image").Select("device_container.id, device_container.created_on, device_container.modified_on, device_container.deleted_on, device_container.container_name, device_container.image_id, device_container.status, device_container.active").Joins("left join device_container on device_image.id = device_container.image_id").Where("device_image.device_id = ? AND device_image.deleted_on = ? AND device_container.deleted_on = ?", device.ID, 0, 0).Scan(&listContainers).Error
	if err != nil {
		logging.Warn(err)
		return listContainers, err
	}
	return listContainers, nil
}

func StopContainer(containerID int) error {
	var deviceContainer DeviceContainer
	err := db.Model(&deviceContainer).Where("id = ? AND deleted_on = ? ", containerID, 0).Updates(
		DeviceContainer{
			Active: "stopping",
		}).Error
	if err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func StartContainer(containerID int) error {
	var deviceContainer DeviceContainer
	err := db.Model(&deviceContainer).Where("id = ? AND deleted_on = ? ", containerID, 0).Updates(
		DeviceContainer{
			Active: "starting",
		}).Error
	if err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func StopAllContainer(deviceID int) error {
	// var deviceContainer DeviceContainer
	// err := db.Model(&deviceContainer).Where("id = ? AND deleted_on = ? ", containerID, 0).Updates(
	// 	DeviceContainer{
	// 		Active: "stopping",
	// 	}).Error
	// if err != nil {
	// 	logging.Warn(err)
	// 	return err
	// }
	return nil
}

func GetMachineIDByContainerID(containerID int) (string, error) {
	var deviceContainer DeviceContainer
	var deviceImage DeviceImage
	var device Device
	db.Where("id = ? AND deleted_on = ?", containerID, 0).First(&deviceContainer)
	db.Where("id = ? AND deleted_on = ?", deviceContainer.ImageID, 0).First(&deviceImage)
	err := db.Where("id = ? AND deleted_on = ?", deviceImage.DeviceID, 0).First(&device).Error

	if err != nil {
		logging.Warn(err)
		return "", err
	}
	return device.MachineID, nil
}

func GetContainerStop(containerID int) (DeviceContainer, error) {
	var deviceContainer DeviceContainer
	err := db.Where("id = ? AND deleted_on = ?", containerID, 0).First(&deviceContainer).Error

	if err != nil {
		logging.Warn(err)
		return deviceContainer, err
	}
	return deviceContainer, nil
}

func CheckStatusStartStop(containerID int) (DeviceContainer, error) {
	var deviceContainer DeviceContainer
	err := db.Where("id = ? AND deleted_on = ?", containerID, 0).First(&deviceContainer).Error
	if err != nil {
		logging.Warn(err)
		return deviceContainer, err
	}
	if deviceContainer.Active == "starting" || deviceContainer.Active == "stopping" {
		logging.Warn(errors.New("Status fail"))
		return deviceContainer, errors.New("Status fail")
	}
	return deviceContainer, nil
}

func UpdateStatusContainer(id int, active string) (DeviceContainer, error) {
	var deviceContainer DeviceContainer
	err := db.Model(&deviceContainer).Where("id = ? AND deleted_on = ? ", id, 0).Updates(
		DeviceContainer{
			Active: active,
		}).Error
	if err != nil {
		logging.Warn(err)
		return deviceContainer, err
	}

	return deviceContainer, nil
}
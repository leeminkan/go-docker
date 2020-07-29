package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type Device struct {
	Model

	DeviceName string `json:"device_name"`
	OS         string `json:"os"`
	MachineID  string `json:"machine_id"`
}

func CreateDevice(device_name string, os string, machine_id string) error {
	device := Device{
		DeviceName: device_name,
		OS:         os,
		MachineID:  machine_id,
	}
	if err := db.Create(&device).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func GetListDevices() ([]Device, error) {
	devices := []Device{}
	result := db.Where("deleted_on = ?", 0).Find(&devices)

	if result.Error != nil {
		logging.Warn(result.Error)
		return nil, result.Error
	}

	return devices, nil
}

func DeleteDevice(id int) error {

	if err := db.Where("id = ?", id).Delete(&Device{}).Error; err != nil {
		logging.Warn(err)
		return err
	}

	return nil
}

func ExistDeviceByID(id int) (bool, error) {
	var device Device
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&device).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if device.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetMachineID(id int) (string, error) {
	logging.Warn(id)
	var device Device
	err := db.Select("machine_id").Where("id = ? AND deleted_on = ? ", id, 0).First(&device).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		logging.Warn(err)
		return "", err
	}
	return device.MachineID, nil
}

func FindDeviceByMachineID(id string) (bool, Device, error) {
	var device Device
	err := db.Select("id").Where("machine_id = ? AND deleted_on = ? ", id, 0).First(&device).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, device, err
	}

	if device.ID > 0 {
		return true, device, nil
	}

	return false, device, nil
}

func (d *Device) Update(device_name string, os string, machine_id string) error {
	err := db.Model(&d).Updates(Device{DeviceName: device_name, OS: os, MachineID: machine_id}).Error
	if err != nil {
		return err
	}

	return nil
}

func CheckDevice(deviceId int) (string, bool, error) {
	var device Device
	err := db.Where("deleted_on = ? AND id = ? ", 0, deviceId).First(&device).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", false, err
	}
	if device.ID > 0 {
		return device.MachineID, true, nil
	}

	return "", false, nil
}

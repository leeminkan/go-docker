package device_service

import (
	"go-docker/models"
	"go-docker/pkg/logging"
)

type Device struct {
	ID         int
	DeviceName string
	OS         string
	MachineID  string
}

func (d *Device) Create() error {
	return models.CreateDevice(d.DeviceName, d.OS, d.MachineID)
}

func GetList() ([]models.Device, error) {
	devices, err := models.GetListDevices()

	if err != nil {
		logging.Warn(err)
		return nil, err
	}
	return devices, nil
}

func (d *Device) Delete() error {
	return models.DeleteDevice(d.ID)
}

func (t *Device) ExistByID() (bool, error) {
	return models.ExistDeviceByID(t.ID)
}

func GetMachineIDFromID(id int) (string, error) {
	return models.GetMachineID(id)
}

func (t *Device) FindByMachineID() (bool, models.Device, error) {
	return models.FindDeviceByMachineID(t.MachineID)
}

func (t *Device) GetListImagesByID() ([]models.DeviceImage, error) {
	return models.GetListImages(t.ID)
}

func (t *Device) GetListContainersByID() ([]models.DeviceContainer, error) {
	return models.GetListContainers(t.ID)
}

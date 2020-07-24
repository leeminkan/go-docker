package device

type CreateDeviceInput struct {
	DeviceName string `form:"device_name" valid:"Required"`
	OS         string `form:"os" valid:"Required"`
	MachineID  string `form:"machine_id" valid:"Required"`
}

type ControlDevicePull struct {
	CreateDeviceInput
	RepoName string `form:"repo_name" valid:"Required"`
}

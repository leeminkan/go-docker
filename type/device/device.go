package device

type CreateDeviceInput struct {
	DeviceName string `form:"device_name" valid:"Required"`
	OS         string `form:"os" valid:"Required"`
	MachineID  string `form:"machine_id" valid:"Required"`
	Status     string `form:"status" valid:"Required"`
}

type ControlDevicePull struct {
	DeviceID int `form:"device_id" valid:"Required"`
	RepoID   int `form:"repo_id" valid:"Required"`
}

type UpdateStatusDevicePull struct {
	FullRepoName string `form:"full_repo_name" valid:"Required"`
	MachineID    string `form:"machine_id" valid:"Required"`
	ImageID      string `form:"image_id"`
	Status       string `form:"status" valid:"Required"`
}

type ControlDeviceRun struct {
	ImagePullID   int    `form:"imagepull_id" valid:"Required"`
	ContainerName string `form:"container_name" valid:"Required"`
}

type UpdateStatusDeviceRun struct {
	ContainerName string `form:"container_name" valid:"Required"`
	ImagePullID   string `form:"imagepull_id" valid:"Required"`
	Status        string `form:"status" valid:"Required"`
	Active        string `form:"active" valid:"Required"`
}

type StopContainer struct {
	ContainerID int `form:"container_id" valid:"Required"`
}

type StartContainer struct {
	ContainerID int `form:"container_id" valid:"Required"`
}

type StopAllContainer struct {
	DeviceID int `form:"device_id" valid:"Required"`
}

type UpdateStatusContainer struct {
	ContainerID string `form:"container_id" valid:"Required"`
	Active      string `form:"active" valid:"Required"`
}

type DeleteContainer struct {
	ContainerID int `form:"container_id" valid:"Required"`
}

type UpdateDeleteContainer struct {
	ContainerID string `form:"container_id" valid:"Required"`
	DeletedOn   string `form:"delete_on" valid:"Required"`
}

type DeleteImage struct {
	ImageID int `form:"image_id" valid:"Required"`
}

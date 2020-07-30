package v1

import (
	"net/http"
	"strconv"

	"go-docker/pkg/app"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/pkg/mqtt"
	"go-docker/service/device_service"
	deviceType "go-docker/type/device"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var GlobalClient = mqtt.InitMQTT()

// @Summary Create devices
// @Produce  json
// @Accept  application/json
// @Tags  Devices
// @Param body body device.CreateDeviceInput true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /devices [post]
func CreateDevice(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.CreateDeviceInput
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	deviceService := device_service.Device{
		DeviceName: form.DeviceName,
		OS:         form.OS,
		MachineID:  form.MachineID,
	}

	exists, _, err := deviceService.FindByMachineID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_DEVICE_FAIL, nil)
		return
	}

	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_DEVICE, nil)
		return
	}

	err = deviceService.Create()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_DEVICE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Get list devices
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Devices
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /devices [get]
func GetListDevices(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	devices, err := device_service.GetList()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LIST_DEVICE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, devices)
}

// @Summary Remove device
// @Produce  json
// @Tags  Devices
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /devices/{id} [delete]
func RemoveDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID is invalid!")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	deviceService := device_service.Device{ID: id}
	exists, err := deviceService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_DEVICE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_DEVICE, nil)
		return
	}

	err = deviceService.Delete()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_DEVICE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "Successful!")
}

// @Summary Device connect
// @Produce  json
// @Accept  application/json
// @Tags  Devices
// @Param body body device.CreateDeviceInput true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /devices/connect [post]
func ConnectDevice(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.CreateDeviceInput
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		logging.Warn(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	deviceService := device_service.Device{
		DeviceName: form.DeviceName,
		OS:         form.OS,
		MachineID:  form.MachineID,
		Status:     form.Status,
	}

	exists, device, err := deviceService.FindByMachineID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CONNECT_DEVICE_FAIL, nil)
		return
	}

	if exists {
		err = device.Update(deviceService.DeviceName, deviceService.OS, deviceService.MachineID, deviceService.Status)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_CONNECT_DEVICE_FAIL, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, device)
		return
	}

	err = deviceService.Create()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CONNECT_DEVICE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Control Device pull image
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.ControlDevicePull true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /control/devices/pull [post]
func ControlDevicePull(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.ControlDevicePull
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	connected, errConnected := device_service.CheckDeviceConnected(form.DeviceID)
	if errConnected != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_SQL_CHECK_CONNECTED, nil)
		return
	}
	if !connected {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CHECK_CONNECTED, nil)
		return
	}

	deviceService := device_service.Control{
		DeviceID: form.DeviceID,
		RepoID:   form.RepoID,
	}

	machineId, exists, err := deviceService.CheckDevice()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_DEVICE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_DEVICE_CONTROL_FAIL, nil)
		return
	}

	repoName, err := deviceService.GetFullRepoNameFromID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_REPONAME_CONTROL, nil)
		return
	}

	deviceImage, errCreatePull := deviceService.CreateDevicePull()
	if errCreatePull != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_UPDATE_STATUS_PULL_FIRST, nil)
		return
	}

	value, err := mqtt.SetValueComeinandPull(machineId, repoName)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
		return
	}
	if value != nil {
		tokenPub := GlobalClient.Publish("image/pull", 0, false, value)
		tokenPub.Wait()
	}

	appG.Response(http.StatusOK, e.SUCCESS, deviceImage)
}

// @Summary Update status image pull from devices
// @Produce  json
// @Accept  application/json
// @Tags  Devices
// @Param body body device.UpdateStatusDevicePull true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/devices/pull/status [post]
func UpdateStatusImagePull(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.UpdateStatusDevicePull
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	updateDeviceService := device_service.Update{
		FullRepoName: form.FullRepoName,
		MachineID:    form.MachineID,
		ImageID:      form.ImageID,
		Status:       form.Status,
	}

	deviceImage, err := updateDeviceService.UpdateImagePullStatus()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_UPDATE_STATUS_IMAGE_PULL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, deviceImage)
}

// @Summary Control Device run image as a container
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.ControlDeviceRun true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /control/devices/run [post]
func ControlDeviceRun(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.ControlDeviceRun
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	deviceService := device_service.ControlRun{
		ImagePullID:   form.ImagePullID,
		ContainerName: form.ContainerName,
	}

	deviceImage, exists, err := deviceService.CheckValueRun()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_DEVICE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_DEVICE_CONTROL_FAIL, nil)
		return
	}

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_REPONAME_CONTROL, nil)
		return
	}

	connected, errConnected := device_service.CheckDeviceConnected(deviceImage.DeviceID)
	if errConnected != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_SQL_CHECK_CONNECTED, nil)
		return
	}
	if !connected {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CHECK_CONNECTED, nil)
		return
	}

	machineID, err := device_service.GetMachineIDFromID(deviceImage.DeviceID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_GET_MACHINE_ID_FROM_ID, nil)
		return
	}

	deviceContainer, err := deviceService.CreateDeviceRun()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CONTAINER_RUN_FIRST, nil)
		return
	}

	value, err := mqtt.SetValueComeinandRun(form.ContainerName, machineID, deviceImage.FullRepoName, form.ImagePullID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
		return
	}
	if value != nil {
		tokenPub := GlobalClient.Publish("container/run", 0, false, value)
		tokenPub.Wait()
	}

	appG.Response(http.StatusOK, e.SUCCESS, deviceContainer)
}

// @Summary Update status container run from devices
// @Produce  json
// @Accept  application/json
// @Tags  Devices
// @Param body body device.UpdateStatusDeviceRun true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/devices/run [post]
func UpdateStatusContainerRun(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.UpdateStatusDeviceRun
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	updateDeviceService := device_service.UpdateRun{
		ContainerName: form.ContainerName,
		ImagePullID:   form.ImagePullID,
		Status:        form.Status,
		Active:        form.Active,
	}

	deviceContainer, err := updateDeviceService.UpdateContainerRunStatus()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_UPDATE_STATUS_IMAGE_PULL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, deviceContainer)
}

// @Summary Get List Image in device
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Devices
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/images/{id} [get]
func GetImagesDeviceByID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()

	connected, errConnected := device_service.CheckDeviceConnected(id)
	if errConnected != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_SQL_CHECK_CONNECTED, nil)
		return
	}
	if !connected {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CHECK_CONNECTED, nil)
		return
	}

	deviceService := device_service.Device{
		ID: id,
	}

	listImages, err := deviceService.GetListImagesByID()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_LIST_IMAGES, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, listImages)
	return
}

// @Summary Get List Container Running in device
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Devices
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/containers/{id} [get]
func GetContainersDeviceByID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()

	connected, errConnected := device_service.CheckDeviceConnected(id)
	if errConnected != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_SQL_CHECK_CONNECTED, nil)
		return
	}
	if !connected {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CHECK_CONNECTED, nil)
		return
	}

	deviceService := device_service.Device{
		ID: id,
	}

	listImages, err := deviceService.GetListContainersByID()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_LIST_IMAGES, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, listImages)
	return
}

// @Summary Get Image Pulling to watch status in device
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Devices
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/image/{id} [get]
func GetImageDeviceByID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()

	deviceService := device_service.StatusPull{
		ImagePullID: id,
	}

	image, err := deviceService.GetImagePull()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_IMAGE_STATUS, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, image)
	return
}

// @Summary Get Container to watch status in device
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Devices
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/container/{id} [get]
func GetContainerDeviceByID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()

	deviceService := device_service.StatusRun{
		ContainerID: id,
	}

	container, err := deviceService.GetContainerRun()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_CONTAINER_STATUS, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, container)
	return
}

// @Summary Stop a Container
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.StopContainer true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/stop/container [post]
func StopContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.StopContainer
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	_, errc := device_service.CheckStatusStartStop(form.ContainerID)
	if errc != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_STATUS_START_STOP_CONTAINER, nil)
		return
	}

	deviceService := device_service.StopContainer{
		ContainerID: form.ContainerID,
	}

	err := deviceService.StopContainerByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_STOP_CONTAINER, nil)
		return
	}

	machineID, err := deviceService.GetMachineIDByContainerID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_GET_MACHINE_ID_FROM_CONTAINER_ID, nil)
		return
	}

	//control device stop container throw mqtt
	container, err := device_service.GetContainer(form.ContainerID)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_CONTAINER_FAIL, nil)
		return
	}

	value, err := mqtt.SetValueComeinandStopContainer(machineID, container.ContainerName, form.ContainerID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
		return
	}
	if value != nil {
		tokenPub := GlobalClient.Publish("container/stop", 0, false, value)
		tokenPub.Wait()
	}

	appG.Response(http.StatusOK, e.SUCCESS, container)
}

// @Summary Start a Container
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.StartContainer true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/start/container [post]
func StartContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.StartContainer
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	_, errc := device_service.CheckStatusStartStop(form.ContainerID)
	if errc != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_STATUS_START_STOP_CONTAINER, nil)
		return
	}

	deviceService := device_service.StartContainer{
		ContainerID: form.ContainerID,
	}

	err := deviceService.StartContainerByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_START_CONTAINER, nil)
		return
	}

	machineID, err := deviceService.GetMachineIDStop()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_GET_MACHINE_ID_FROM_CONTAINER_ID, nil)
		return
	}

	//control device stop container throw mqtt
	container, err := device_service.GetContainerStart(form.ContainerID)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_CONTAINER_FAIL, nil)
		return
	}

	value, err := mqtt.SetValueComeinandStopContainer(machineID, container.ContainerName, form.ContainerID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
		return
	}
	if value != nil {
		tokenPub := GlobalClient.Publish("container/start", 0, false, value)
		tokenPub.Wait()
	}

	appG.Response(http.StatusOK, e.SUCCESS, container)
}

// @Summary Stop all Container
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.StopAllContainer true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/stop/container/all [post]
func StopAllContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.StopAllContainer
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	deviceService := device_service.StopContainer{
		DeviceID: form.DeviceID,
	}

	err := deviceService.StopContainerByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_STOP_CONTAINER, nil)
		return
	}

	machineID, err := deviceService.GetMachineIDByContainerID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_GET_MACHINE_ID_FROM_CONTAINER_ID, nil)
		return
	}

	//control device stop container throw mqtt
	// deviceServiceContainer := device_service.StatusRun{
	// 	ContainerID: form.ContainerID,
	// }
	// container, err := deviceServiceContainer.GetContainerRun()

	// if err != nil {
	// 	appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_CONTAINER_STATUS, nil)
	// 	return
	// }

	// value, err := mqtt.SetValueComeinandStopContainer()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
	// 	return
	// }
	// if value != nil {
	// 	tokenPub := GlobalClient.Publish("container/run", 0, false, value)
	// 	tokenPub.Wait()
	// }

	appG.Response(http.StatusOK, e.SUCCESS, machineID)
}

// @Summary Start a Container
// @Produce  json
// @Accept  application/json
// @Tags  Devices
// @Param body body device.UpdateStatusContainer true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/update/container/status [post]
func UpdateStatusContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.UpdateStatusContainer
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	cvtInt, _ := strconv.Atoi(form.ContainerID)

	deviceService := device_service.StatusRun{
		ContainerID: cvtInt,
		Active:      form.Active,
	}

	container, err := deviceService.UpdateStatusContainer()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_UPDATE_STATUS_START_STOP_CONTAINER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, container)
}

// @Summary Delete a Container
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.DeleteContainer true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/delete/container [post]
func DeleteContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.DeleteContainer
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	delete, errc := device_service.CheckDeleteContainer(form.ContainerID)
	if errc != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_SQL_DELETE_CONTAINER, nil)
		return
	}
	if !delete {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CHECK_DELETE_CONTAINER, nil)
		return
	}

	deviceService := device_service.StopContainer{
		ContainerID: form.ContainerID,
	}

	machineID, err := deviceService.GetMachineIDByContainerID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_GET_MACHINE_ID_FROM_CONTAINER_ID, nil)
		return
	}

	//control device stop container throw mqtt
	container, err := device_service.GetContainer(form.ContainerID)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_CONTAINER_FAIL, nil)
		return
	}

	value, err := mqtt.SetValueComeinandRemoveContainer(machineID, container.ContainerName, form.ContainerID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
		return
	}
	if value != nil {
		tokenPub := GlobalClient.Publish("container/delete", 0, false, value)
		tokenPub.Wait()
	}

	appG.Response(http.StatusOK, e.SUCCESS, container)
}

// @Summary Update status delete a Container
// @Produce  json
// @Accept  application/json
// @Tags  Devices
// @Param body body device.UpdateDeleteContainer true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/update/container/delete [post]
func UpdateDeleteContainer(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.UpdateDeleteContainer
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	intContainerID, _ := strconv.Atoi(form.ContainerID)
	intDeleteOn, _ := strconv.Atoi(form.DeletedOn)

	deviceService := device_service.DeleteContainer{
		ContainerID: intContainerID,
		DeleteOn:    intDeleteOn,
	}

	container, err := deviceService.UpdateDeleteContainer()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_UPDATE_STATUS_START_STOP_CONTAINER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, container)
}

// @Summary Delete a Image
// @Produce  json
// @Security ApiKeyAuth
// @Accept  application/json
// @Tags  Devices
// @Param body body device.DeleteImage true "body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /device/delete/image [post]
func DeleteImage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form deviceType.DeleteImage
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	delete, errc := device_service.CheckDeleteImage(form.ImageID)
	if errc != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_SQL_DELETE_IMAGE, nil)
		return
	}
	if !delete {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_CHECK_DELETE_IMAGE, nil)
		return
	}

	deviceService := device_service.DeleteImage{
		ImageID: form.ImageID,
	}

	machineID, err := deviceService.GetMachineIDByImageID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DEVICE_GET_MACHINE_ID_FROM_IMAGE_ID, nil)
		return
	}

	// //control device stop container throw mqtt
	// container, err := device_service.Get

	// if err != nil {
	// 	appG.Response(http.StatusBadRequest, e.ERROR_DEVICE_GET_CONTAINER_FAIL, nil)
	// 	return
	// }

	// value, err := mqtt.SetValueComeinandRemoveContainer(machineID, container.ContainerName, form.ContainerID)
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
	// 	return
	// }
	// if value != nil {
	// 	tokenPub := GlobalClient.Publish("container/delete", 0, false, value)
	// 	tokenPub.Wait()
	// }

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

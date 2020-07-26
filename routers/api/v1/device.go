package v1

import (
	"net/http"

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
		appG.Response(httpCode, errCode, nil)
		return
	}

	deviceService := device_service.Device{
		DeviceName: form.DeviceName,
		OS:         form.OS,
		MachineID:  form.MachineID,
	}

	exists, device, err := deviceService.FindByMachineID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CONNECT_DEVICE_FAIL, nil)
		return
	}

	if exists {
		err = device.Update(deviceService.DeviceName, deviceService.OS, deviceService.MachineID)
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

	deviceService := device_service.Control{
		DeviceName: form.DeviceName,
		OS:         form.OS,
		MachineID:  form.MachineID,
		RepoID:     form.RepoID,
	}

	exists, err := deviceService.ExistDevice()
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

	value, err := mqtt.SetValueComeinand(form.MachineID, repoName)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SET_MESSAGE_MQTT, nil)
		return
	}
	if value != nil {
		tokenPub := GlobalClient.Publish("image/pull", 0, false, value)
		tokenPub.Wait()
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

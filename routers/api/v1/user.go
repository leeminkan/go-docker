package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-docker/models"
	"go-docker/pkg/app"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/pkg/util"
	"go-docker/service/user_service"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username        string `form:"username" valid:"Required"`
	Password        string `form:"password" valid:"Required"`
	ConfirmPassword string `form:"confirm_password" valid:"Required"`
	IsAdmin         bool   `form:"is_admin" valid:"Required"`
}

// @Summary Create users
// @Produce  json
// @Accept  multipart/form-data
// @Tags  Users
// @Param username formData string true "userName"
// @Param password formData string true "password"
// @Param confirm_password formData string true "confirm_password"
// @Param is_admin formData boolean true "is_admin"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form CreateUserInput
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	if form.Password != form.ConfirmPassword {
		appG.Response(http.StatusBadRequest, e.ERROR_MATCH_CONFIRM_PASSWORD_USER_FAIL, nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER_FAIL, nil)
		return
	}

	userService := user_service.User{
		Username: form.Username,
		Password: string(hashedPassword),
		IsAdmin:  form.IsAdmin,
	}

	exists, err := userService.ExistByUserName()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER_FAIL, nil)
		return
	}

	if exists {
		appG.Response(http.StatusConflict, e.ERROR_EXIST_USER_FAIL, nil)
		return
	}

	err = userService.Create()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type LoginUserInput struct {
	Username string `form:"username" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

// @Summary User Login
// @Produce  json
// @Accept  multipart/form-data
// @Tags  Users
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/login [post]
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form LoginUserInput
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		logging.Warn(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
	}
	token, err := userService.Login()

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_LOGIN_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
	return
}

// @Summary Get Info User
// @Produce  json
// @Security ApiKeyAuth
// @Tags  Users
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/mine [get]
func GetInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	user, _ := c.MustGet("user").(models.User)
	data, _ := util.DecodeBase64XRegistryAuth(user.XRegistryAuth)

	appG.Response(http.StatusOK, e.SUCCESS, data.Username)
	return
}

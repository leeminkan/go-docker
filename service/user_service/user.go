package user_service

import (
	"go-docker/models"
	"go-docker/pkg/logging"
	"go-docker/pkg/util"
)

type User struct {
	ID       int
	Username string
	Password string
	IsAdmin  bool
}

func (u *User) Create() error {
	return models.CreateUser(u.Username, u.Password, u.IsAdmin)
}

func (u *User) ExistByUserName() (bool, error) {
	return models.ExistByUserName(u.Username)
}

func (u *User) Login() (string, error) {
	err := models.CheckLogin(u.Username, u.Password)
	if err != nil {
		logging.Warn(err)
		return "", err
	}

	token, err := util.GenerateToken(u.Username)
	if err != nil {
		logging.Warn(err)
		return "", err
	}

	return token, nil
}

func GetUserByUserName(username string) (models.User, error) {
	user, err := models.GetUserByUserName(username)
	if err != nil {
		logging.Warn(err)
		return user, err
	}

	return user, nil
}

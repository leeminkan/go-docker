package models

import (
	"go-docker/pkg/logging"

	"go-docker/pkg/e"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model

	Username         string `json:"username"`
	Password         string `json:"password"`
	IsAdmin          bool   `json:"is_admin"`
	XRegistryAuth    string `json:"x_registry_auth"`
	IsLoginDockerHub bool   `json:"is_login_docker_hub"`
}

func CreateUser(username string, password string, is_admin bool) error {
	user := User{
		Username: username,
		Password: password,
		IsAdmin:  is_admin,
	}
	if err := db.Create(&user).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

// ExistByUserName checks if there is a user with the same name
func ExistByUserName(username string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CheckLogin(username, password string) error {
	var user User
	err := db.Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return err
	}

	logging.Warn(user)
	if user.ID > 0 {

		// Comparing the password with the hash
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			logging.Warn(err)
			return err
		}

		return nil
	}

	return e.New("User not found!")
}

// GetUserByUserName
func GetUserByUserName(username string) (User, error) {
	var user User
	err := db.Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return user, err
	}

	if user.ID > 0 {
		return user, nil
	}

	return user, nil
}

// update X-Registry-Auth
func (user User) UpdateXRegistryAuth(login bool, xRegistryAuth string) error {
	if login == true {
		err := db.Model(&user).Where("deleted_on = ?", 0).Updates(map[string]interface{}{"x_registry_auth": xRegistryAuth, "is_login_docker_hub": true}).Error

		if err != nil {
			logging.Warn(err)
			return err
		}
	} else {
		err := db.Model(&user).Where("deleted_on = ?", 0).Updates(map[string]interface{}{"x_registry_auth": nil, "is_login_docker_hub": false}).Error

		if err != nil {
			logging.Warn(err)
			return err
		}
	}

	return nil
}

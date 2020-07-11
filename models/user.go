package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
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

	return nil
}

// GetUserByUserName
func GetUserByUserName(username string) (User, error) {
	var user User
	err := db.Select("id").Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return user, err
	}

	if user.ID > 0 {
		return user, nil
	}

	return user, nil
}

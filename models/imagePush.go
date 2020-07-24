package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type ImagePush struct {
	Model

	RepoName    string `json:"repo_name"`
	UserID      int    `json:"user_id"`
	Status      string `json:"status"`
	OldRepoName string `json:"old_repo_name"`
}

func CreateImagePush(repo_name string, user_id int, status string) (ImagePush, error) {
	imagePush := ImagePush{
		RepoName: repo_name,
		UserID:   user_id,
		Status:   status,
	}

	if err := db.Create(&imagePush).Error; err != nil {
		logging.Warn(err)
		return imagePush, err
	}

	return imagePush, nil
}

func UpdateStatusImagePush(repo_name string, status string) (ImagePush, error) {
	imagePush := ImagePush{
		RepoName: repo_name,
		Status:   status,
	}
	db.Table("image_push").Where("repo_name = ? AND deleted_on = ? ", repo_name, 0).Update("status", status)
	err := db.Where("deleted_on = ?", 0).Find(&imagePush).Error
	if err != nil {
		logging.Warn(err)
		return imagePush, err
	}

	return imagePush, nil
}

func GetListImagePush() ([]ImagePush, error) {
	var images []ImagePush
	err := db.Where("deleted_on = ?", 0).Find(&images).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return images, err
	}

	return images, nil
}

func CheckExistRepoToRefuse(repo_name string) bool {
	var image ImagePush
	db.Where("repo_name = ? AND deleted_on = ? ", repo_name, 0).Find(&image)
	if image == (ImagePush{}) {
		return false
	}
	return true
}

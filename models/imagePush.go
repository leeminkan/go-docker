package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type ImagePush struct {
	Model

	TagID        int    `json:"tag_id"`
	UserID       int    `json:"user_id"`
	BuildID      int    `json:"build_id"`
	FullRepoName string `json:"full_repo_name"`
	Digest       string `json:"digest"`
	Status       string `json:"status" gorm:"type:enum('on progress', 'done', 'fail');default:'on progress'"`

	TagDockerHub TagDockerHub `gorm:"foreignkey:TagID"`
	User         User
	ImageBuild   ImageBuild `gorm:"foreignkey:BuildID"`
}

func CreateImagePush(tag_id int, user_id int, build_id int, full_repo_name string, digest string, status string) (ImagePush, error) {
	imagePush := ImagePush{
		TagID:        tag_id,
		UserID:       user_id,
		BuildID:      build_id,
		FullRepoName: full_repo_name,
		Digest:       digest,
		Status:       status,
	}

	if err := db.Create(&imagePush).Error; err != nil {
		logging.Warn(err)
		return imagePush, err
	}

	db.Preload("TagDockerHub").Preload("TagDockerHub.RepoDockerHub").Preload("User").Preload("ImageBuild").Find(&imagePush)

	return imagePush, nil
}

func GetListImagePush() ([]ImagePush, error) {
	var images []ImagePush
	err := db.Where("deleted_on = ?", 0).Preload("TagDockerHub").Preload("TagDockerHub.RepoDockerHub").Preload("User").Preload("ImageBuild").Find(&images).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return images, err
	}

	return images, nil
}

func GetImagePushByID(id int) (bool, ImagePush, error) {
	var image ImagePush
	err := db.Where("id = ? AND deleted_on = ?", id, 0).Preload("TagDockerHub").Preload("TagDockerHub.RepoDockerHub").Preload("User").Preload("ImageBuild").First(&image).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, image, err
	}

	if image.ID > 0 {
		return true, image, nil
	}

	return false, image, nil
}

func (image ImagePush) UpdatePush(digest string, status string) error {
	err := db.Model(&image).Updates(ImagePush{
		Digest: digest,
		Status: status}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetFullRepoName(pushID int) (string, error) {
	var imagePush ImagePush

	err := db.Select("full_repo_name").Where("deleted_on = ? AND id = ? ", 0, pushID).First(&imagePush).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if imagePush.FullRepoName != "" {
		return imagePush.FullRepoName, nil
	}

	return "", nil
}

func CheckImagePushExist(tag_id int, build_id int, statusDone string, statusOnProgress string) (bool, ImagePush) {
	var image ImagePush
	err := db.Where("tag_id = ? AND build_id = ? AND (status = ? OR status = ?) AND deleted_on = ? ", tag_id, build_id, statusDone, statusOnProgress, 0).First(&image).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, image
	}

	if image.ID > 0 {
		return true, image
	}

	return false, image
}

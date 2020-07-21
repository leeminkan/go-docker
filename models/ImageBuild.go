package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type ImageBuild struct {
	Model

	RepoName    string `json:"repo_name"`
	ImageID     string `json:"image_id"`
	UserID      int    `json:"user_id"`
	Status      string `json:"status"`
	OldRepoName string `json:"old_repo_name"`
}

func CreateImageBuild(repo_name string, image_id string, user_id int, status string) (ImageBuild, error) {
	imageBuild := ImageBuild{
		RepoName: repo_name,
		ImageID:  image_id,
		UserID:   user_id,
		Status:   status,
	}

	if err := db.Create(&imageBuild).Error; err != nil {
		logging.Warn(err)
		return imageBuild, err
	}

	return imageBuild, nil
}

func RemoveRepoNameIfExist(repo_name string) error {
	var image ImageBuild
	err := db.Model(&image).Where("repo_name = ? AND deleted_on = ? ", repo_name, 0).Updates(map[string]interface{}{"repo_name": "", "old_repo_name": repo_name}).Error

	if err != nil {
		logging.Warn(err)
		return err
	}

	return nil
}

func (image ImageBuild) Update(repo_name string, image_id string, user_id int, status string) error {
	err := db.Model(&image).Updates(ImageBuild{
		RepoName: repo_name,
		ImageID:  image_id,
		UserID:   user_id,
		Status:   status}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetImageBuildByID(id int) (bool, ImageBuild, error) {
	var image ImageBuild
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&image).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, image, err
	}

	if image.ID > 0 {
		return true, image, nil
	}

	return false, image, nil
}

func GetImageBuild(repo_name string, image_id string) (bool, ImageBuild, error) {
	var image ImageBuild
	err := db.Where("repo_name = ? AND deleted_on = ? AND image_id = ?", repo_name, 0, image_id).First(&image).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, image, err
	}

	if image.ID > 0 {
		return true, image, nil
	}

	return false, image, nil
}

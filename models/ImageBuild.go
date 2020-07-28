package models

import (
	"go-docker/pkg/logging"

	"github.com/jinzhu/gorm"
)

type ImageBuild struct {
	Model

	RepoName    string `json:"repo_name"`
	Tag         string `json:"tag" gorm:"default:'latest'"`
	ImageID     string `json:"image_id"`
	UserID      int    `json:"user_id"`
	Status      string `json:"status" gorm:"type:enum('on progress', 'done', 'fail');default:'on progress'"`
	OldRepoName string `json:"old_repo_name"`
}

type ImageBuildCustom struct {
	ImageBuild
	IsPushed int `json:"is_pushed"`
}

func GetListImageBuildCustom(user_id int) ([]ImageBuildCustom, error) {
	var images []ImageBuildCustom
	rows, err := db.Model(&ImageBuild{}).Where("image_build.deleted_on = ?", 0).Joins("left join (select * from image_push where image_push.deleted_on = ? and image_push.status = ? and image_push.user_id = ?) as image_push on image_push.build_id = image_build.id", 0, "done", user_id).Select("image_build.*, COUNT(image_push.id) is_pushed").Group("image_build.id").Rows()
	if err != nil {
		logging.Warn(err)
		return images, err
	}
	defer rows.Close()
	for rows.Next() {
		var image ImageBuildCustom
		db.ScanRows(rows, &image)
		images = append(images, image)
	}
	return images, nil
}

func CreateImageBuild(repo_name string, tag string, image_id string, user_id int, status string, old_repo_name string) (ImageBuild, error) {
	imageBuild := ImageBuild{
		RepoName:    repo_name,
		Tag:         tag,
		ImageID:     image_id,
		UserID:      user_id,
		Status:      status,
		OldRepoName: old_repo_name,
	}

	if err := db.Create(&imageBuild).Error; err != nil {
		logging.Warn(err)
		return imageBuild, err
	}

	return imageBuild, nil
}

func RemoveRepoNameAndTagIfExist(repo_name string, tag string) error {
	var image ImageBuild
	err := db.Model(&image).Where("repo_name = ? AND tag = ? AND deleted_on = ? ", repo_name, tag, 0).Updates(map[string]interface{}{"repo_name": "", "tag": "", "old_repo_name": repo_name + ":" + tag}).Error

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

func GetListImageBuild() ([]ImageBuild, error) {
	var images []ImageBuild
	err := db.Where("deleted_on = ?", 0).Find(&images).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return images, err
	}
	return images, nil
}

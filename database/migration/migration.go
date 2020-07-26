package migration

import (
	"go-docker/models"
	"go-docker/service/migration_service"
)

// Để có thể migrate được, các model được truyền vào phải theo định dạng tại: http://gorm.io/docs/models.html
// Hàm AutoMigrate không thể thay đổi trạng thái của column, hoặc xóa để bảo toàn dữ liệu

func Migrate() {
	db := models.GetDB()
	db.AutoMigrate(&models.User{}, &models.Device{}, &models.ImageBuild{}, &models.Seed{}, &models.Migration{}, &models.ImagePush{}, &models.RepoDockerHub{}, &models.TagDockerHub{})
	MigrateForUser()
}

func CheckMigration(name string) bool {
	migration := migration_service.Migration{
		Name: name,
	}

	exist, err := migration.IsMigrationExistedByName()

	if err != nil {
		return false
	}
	if exist {
		return false
	}

	return true
}

func CreateMigration(name string) bool {
	migration := migration_service.Migration{
		Name: name,
	}

	err := migration.Create()

	if err != nil {
		return false
	}

	return true
}

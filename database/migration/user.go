package migration

import (
	"go-docker/models"
)

func MigrateForUser() {
	DropColumnTest()
}

func DropColumnTest() {
	name := "DropColumnTest-25072020"
	check := CheckMigration(name)
	if check != true {
		return
	}

	db := models.GetDB()

	db.Model(&models.User{}).DropColumn("test")

	CreateMigration(name)
}

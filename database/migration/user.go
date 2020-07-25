package migration

import (
	"go-docker/models"
)

func MigrateForUser() {
	DropColumnTest()
	UpdateUniqueForUserName()
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

func UpdateUniqueForUserName() {
	name := "UpdateUniqueForUserName-25072020"
	check := CheckMigration(name)
	if check != true {
		return
	}

	db := models.GetDB()

	db.Exec("ALTER TABLE user ADD UNIQUE (username);")

	CreateMigration(name)
}

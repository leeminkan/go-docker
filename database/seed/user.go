package seed

import (
	"go-docker/service/user_service"
)

func SeedForUser() {
	CreateAdmin()
}

func CreateAdmin() {
	name := "CreateAdmin-25072020"
	check := CheckSeed(name)
	if check != true {
		return
	}

	users := []user_service.User{
		{
			Username: "admin",
			Password: "123456",
			IsAdmin:  true,
		},
	}

	for _, user := range users {
		user.Create()
	}

	CreateSeed(name)
}

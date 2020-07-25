package seed

import (
	"go-docker/service/user_service"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	users := []user_service.User{
		{
			Username: "admin",
			Password: string(hashedPassword),
			IsAdmin:  true,
		},
	}

	for _, user := range users {
		user.Create()
	}

	CreateSeed(name)
}

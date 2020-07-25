package seed

import "go-docker/service/seed_service"

func Run() {
	SeedForUser()
}

func CheckSeed(name string) bool {
	seed := seed_service.Seed{
		Name: name,
	}

	exist, err := seed.IsSeedExistedByName()

	if err != nil {
		return false
	}
	if exist {
		return false
	}

	return true
}

func CreateSeed(name string) bool {
	seed := seed_service.Seed{
		Name: name,
	}

	err := seed.Create()

	if err != nil {
		return false
	}

	return true
}

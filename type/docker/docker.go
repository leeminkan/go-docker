package docker

type LoginDockerInput struct {
	Username string `form:"username" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

package docker

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-docker/models"
	"go-docker/pkg/logging"
)

func CheckLoginDockerHub() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, _ := c.MustGet("user").(models.User)

		if user.IsLoginDockerHub != true {
			logging.Warn("User %s haven't already login in docker hub!", user.Username)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 9999,
				"msg":  "User haven't already login in docker hub!",
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

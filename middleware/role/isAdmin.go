package role

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-docker/models"
	"go-docker/pkg/logging"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, _ := c.MustGet("user").(models.User)

		if user.IsAdmin != true {
			logging.Warn("Permission! ", user.Username)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 9999,
				"msg":  "Permission!",
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

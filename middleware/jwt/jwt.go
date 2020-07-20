package jwt

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go-docker/models"
	"go-docker/pkg/e"
	"go-docker/pkg/logging"
	"go-docker/pkg/util"
	"go-docker/service/user_service"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

// JWT is jwt middleware
func JWTCustom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var user models.User

		code = e.SUCCESS
		token := ExtractToken(c)
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			resolved, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					logging.Warn(err)
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					logging.Warn(err)
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			} else {
				user, err = user_service.GetUserByUserName(resolved.Username)
				if err != nil {
					logging.Warn(err)
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			logging.Warn(code)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

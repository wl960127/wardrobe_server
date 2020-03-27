package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"wardrobe_server/pkg/e"
	"wardrobe_server/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT .
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *utils.Claims
		var err error

		code = e.SUCCESS
		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err = utils.ParseToken(token[1])
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
		fmt.Printf("JWT中间件 %d %s", claims.ID, claims.Mobile)
		c.Set("claimsID", claims.ID)
		c.Set("claimsMobile", claims.Mobile)
		c.Next()
	}
}

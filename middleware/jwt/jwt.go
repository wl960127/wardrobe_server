package jwt

import (
	"wardrobe_server/pkg/msg"
	"net/http"
	"fmt"
	"strings"
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

		code = msg.SUCCESS
		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = msg.INVALID_PARAMS
		} else {
			claims, err = utils.ParseToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = msg.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != msg.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		fmt.Printf("JWT ?? %d %s", claims.ID, claims.Mobile)
		c.Set("claimsID", claims.ID)
		c.Set("claimsMobile", claims.Mobile)
		c.Next()
	}
}

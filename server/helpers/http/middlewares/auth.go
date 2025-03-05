package middlewares

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userAttr         = "userAttr"
	missingHeaderMsg = "Authorization header missing/invalid"
	invalidTokenMsg  = "Invalid token"
)

func Guard() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		fmt.Println("tokenString: ", tokenString)

		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(missingHeaderMsg))
			c.Abort()
			return
		}

		jwtAttr, err := jwt.ValidateToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewUnauthorizedResponse(invalidTokenMsg))
			c.Abort()
			return
		}

		c.Set("userAttr", jwtAttr)
		c.Next()
	}
}

func CheckUserRoles(whitelistRole []string) gin.HandlerFunc {
	mapper := make(map[string]bool)

	for _, s := range whitelistRole {
		mapper[s] = true
	}

	return func(c *gin.Context) {
		value, exists := c.Get(userAttr)
		if !exists {
			c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(missingHeaderMsg))
			c.Abort()
			return
		}

		jwtAttr, ok := value.(jwt.CustomClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, common.NewUnauthorizedResponse(invalidTokenMsg))
			c.Abort()
			return
		}

		if mapper[jwtAttr.RoleID] {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, "Anda tidak punya akses untuk menggunakan fitur ini")
		c.Abort()
	}
}

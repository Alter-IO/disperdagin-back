package jwt

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetJWTAttributes(c *gin.Context) (CustomClaims, error) {
	jwtAttr, exists := c.Get("userAttr")
	if !exists {
		return CustomClaims{}, errors.New("jwt attribute tidak ditemukan dalam konteks")
	}

	userAttr, ok := jwtAttr.(CustomClaims)
	if !ok {
		return userAttr, errors.New("jwt attribute tidak valid")
	}

	return userAttr, nil
}

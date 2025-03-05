package jwt

import (
	cfg "alter-io-go/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(cfg.GetConfig().JWT.Secret)

type CustomClaims struct {
	UserID     string `json:"user_id"`
	RoleID     string `json:"role_id"`
	DivisionID string `json:"division_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(id, roleID, divisionID string) (tokenString string, err error) {
	claims := &CustomClaims{
		UserID:     id,
		RoleID:     roleID,
		DivisionID: divisionID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.GetConfig().JWT.Issuer,
			Subject:   cfg.GetConfig().JWT.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return CustomClaims{}, errors.New("token kadaluarsa")
		}
		return CustomClaims{}, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return CustomClaims{
			UserID:     claims.UserID,
			RoleID:     claims.RoleID,
			DivisionID: claims.DivisionID,
		}, nil
	}

	return CustomClaims{}, errors.New("token tidak valid")
}

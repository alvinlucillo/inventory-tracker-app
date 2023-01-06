package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetJWTKey() []byte {
	jwtString := os.Getenv("inventory-secret")

	if jwtString == "" {
		jwtString = "inventory-secret"
	}

	return []byte(jwtString)
}

func GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(GetJWTKey())

	return
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(GetJWTKey()), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func CheckToken(c *gin.Context) {
	c.Status(http.StatusOK)
}

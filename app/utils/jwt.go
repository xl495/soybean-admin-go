package utils

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateToken(userid uint, userName string) (string, error) {

	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userName"] = userName
	claims["userId"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return t, err
}

package utils

import (
	"distributed_database_server/internal/auth/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// JWT Claims struct
type Claims struct {
	Id       int
	UserName string
	jwt.StandardClaims
}

// Generate new JWT Token
func GenerateJWTToken(user *models.UserResponse, secret string, ttl int) (string, error) {
	claims := Claims{
		Id:       user.Id,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ComparePasswords(hashedPwd string, plainPwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return err
	}
	return nil
}

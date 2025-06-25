package controllers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ для подписи JWT
var jwtSecret = []byte("@Zhandos00")

// Функция для создания JWT-токена
func GenerateJWT(username, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Токен действует 24 часа
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

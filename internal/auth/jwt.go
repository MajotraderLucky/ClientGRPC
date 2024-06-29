package auth

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func init() {
	// Загрузка .env файла при инициализации пакета
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GenerateJWT() (string, error) {
	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Токен истекает через 24 часа
		Issuer:    "exampleIssuer",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "os"
)

func GenerateToken(userId string) (string, error) {
    claims := jwt.MapClaims{
        "id":  userId,
        "exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

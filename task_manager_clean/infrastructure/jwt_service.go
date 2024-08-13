package infrastructure

// Functions to generate and validate JWT tokens.

import (
	"errors"
	"fmt"
	"os"
	"task_manager_clean/domain"

	"github.com/dgrijalva/jwt-go"
)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken generates a new JWT token.
func  GenerateToken(id string, username string, role domain.Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       id,
		"Username": username,
		"Role":     role,
	})

	return token.SignedString(jwtSecret)
}

// ValidateToken validates a JWT token.
func  ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid JWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid JWT claims")
	}

	role, ok := claims["Role"].(string)
	if !ok {
		return "", errors.New("invalid JWT role")
	}

	return role, nil
}

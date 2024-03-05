package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "wouldBeMoreComplexInRealWorldApp"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": strconv.FormatInt(userId, 10),
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // checks it's of the correct type

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token")
	}
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token!")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims!")
	}

	// Check and retrieve userId claim
	userIdClaim, ok := claims["userId"].(string)
	if !ok {
		return 0, errors.New("Invalid 'userId' claim type")
	}

	// Convert userId from string to int64
	userId, err := strconv.ParseInt(userIdClaim, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Failed to convert 'userId' to int64: %v", err)
	}

	return userId, nil
}

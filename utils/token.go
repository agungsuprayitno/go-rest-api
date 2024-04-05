package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(ttl time.Duration, payload string, privateKey string) (string, error) {

	mySigningKey := []byte(privateKey)

	claims := 
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			// NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    payload,
			Subject:   payload,
			// ID:        "1",
			// Audience:  []string{"somebody_else"},
		}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateToken(token string, secretKey string) (interface{}, error) {
	
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims["sub"], nil
}

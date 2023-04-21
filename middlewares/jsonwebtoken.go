package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/config/common"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(common.JWTSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWT(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.JWTSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("invalid or expired token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid or expired token")
	}
	email, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("invalid or expired token")
	}
	return email, nil
}

package middlewares

import (
	"Event-Planning-App/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	// Initialize configuration
	cfg, err := config.InitConfiguration()
	if err != nil {
		panic(err)
	}

	return echoJWT.WithConfig(echoJWT.Config{
		SigningKey:    []byte(cfg.JWTConfig.Secret),
		SigningMethod: cfg.JWTConfig.SigningMethod,
	})
}

func CreateToken(userId uint) (string, error) {
	// Initialize configuration
	cfg, err := config.InitConfiguration()
	if err != nil {
		panic(err)
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTConfig.Secret))
}

func ExtractToken(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}
	return 0
}

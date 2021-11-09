package middleware

import (
	"github.com/golang-jwt/jwt"
	"go-hexagonal-auth/config"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type JwtCustomClaims struct {
	Name  string `json:"name"`
	ID  int `json:"id"`
	Email  string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}


func JWTMiddleware(cfg config.Config ) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		Claims: &JwtCustomClaims{},
		SigningKey:    []byte(cfg.JWTConfig.Secret),
	})
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["isAdmin"].(bool)
		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}


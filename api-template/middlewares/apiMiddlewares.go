package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetApiMiddlewares(group *echo.Group) {
	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS512",
	}))
}

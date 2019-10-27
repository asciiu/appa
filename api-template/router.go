package main

import (
	"database/sql"
	"net/http"

	"github.com/asciiu/appa/api-template/controllers"
	"github.com/asciiu/appa/api-template/middlewares"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// send 200 ok to ping requests
func health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewRouter(db *sql.DB) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	middlewares.SetMainMiddlewares(e)

	//service := k8s.NewService(micro.Name("api"))

	//service.Init()

	// controllers
	authController := controllers.NewAuthController(db)

	// required for health checks
	e.GET("/index.html", health)
	e.GET("/", health)

	// api group
	openApi := e.Group("/api")

	// open endpoints here
	openApi.POST("/login", authController.HandleLogin)
	//openApi.POST("/signup", authController.HandleSignup)

	protectedApi := e.Group("/api")
	// set the auth middlewares
	protectedApi.Use(authController.RefreshAccess)
	middlewares.SetApiMiddlewares(protectedApi)

	// ###########################  protected endpoints here
	protectedApi.GET("/logout", authController.HandleLogout)

	//go func() {
	//	if err := service.Run(); err != nil {
	//		log.Println("nope! ", err)
	//	}
	//}()

	return e
}

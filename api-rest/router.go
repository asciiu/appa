package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	repoQueries "github.com/asciiu/appa/lib/refreshToken/db/sql"
	"github.com/asciiu/appa/api/handlers"
	"github.com/asciiu/appa/api/middlewares"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro/v2"
	k8s "github.com/micro/kubernetes/go/micro"
	"gopkg.in/go-playground/validator.v9"
)

// clean up stage refresh tokens in DB every 30 minutes
const cleanUpInterval = 30 * time.Minute

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// send 200 ok to ping requests
func health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

// routine to clean up refresh tokens in DB
func cleanDatabase(db *sql.DB) {
	for {
		time.Sleep(cleanUpInterval)
		error := repoQueries.DeleteStaleTokens(db, time.Now())
		if error != nil {
			log.Fatal(error)
		}
	}
}

func NewRouter(db *sql.DB) *echo.Echo {
	go cleanDatabase(db)

	e := echo.New()
	//e.AutoTLSManager.Prompt = autocert.AcceptTOS
	//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("stage.appa.exchange")
	//e.AutoTLSManager.Cache = autocert.DirCache("/mnt/appa/autocert")

	middlewares.SetMainMiddlewares(e)

	service := k8s.NewService(micro.Name("api"))

	service.Init()

	// handlers
	authHandler := handlers.NewAuthController(db)
	sessionHandler := handlers.NewSessionController(db)
	userHandler := handlers.NewUserController(db)

	// required for health checks
	e.GET("/index.html", health)
	e.GET("/", health)

	// api group
	openApi := e.Group("/api")

	// open endpoints here
	openApi.POST("/login", authHandler.HandleLogin)
	openApi.POST("/signup", userHandler.HandleSignup)

	protectedApi := e.Group("/api")
	// set the auth middlewares
	protectedApi.Use(authHandler.RefreshAccess)
	protectedApi.Use(authHandler.PopulateContext)
	middlewares.SetApiMiddlewares(protectedApi)

	// ###########################  protected endpoints here
	protectedApi.GET("/session", sessionHandler.HandleSession)
	protectedApi.GET("/logout", authHandler.HandleLogout)

	// user manangement endpoints
	protectedApi.PUT("/users/:userID/changepassword", userHandler.HandleChangePassword)
	protectedApi.PUT("/users/:userID", userHandler.HandleUpdateUser)

	go func() {
		if err := service.Run(); err != nil {
			log.Println("nope! ", err)
		}
	}()

	return e
}

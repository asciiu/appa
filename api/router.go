package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/asciiu/oldiez/api/controllers"
	repoToken "github.com/asciiu/oldiez/api/db/sql"
	"github.com/asciiu/oldiez/api/middlewares"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"golang.org/x/crypto/acme/autocert"
)

// clean up stage refresh tokens in DB every 30 minutes
const cleanUpInterval = 30 * time.Minute

// send 200 ok to ping requests
func health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

// routine to clean up refresh tokens in DB
func cleanDatabase(db *sql.DB) {
	for {
		time.Sleep(cleanUpInterval)
		error := repoToken.DeleteStaleTokens(db, time.Now())
		if error != nil {
			log.Fatal(error)
		}
	}
}

func NewRouter(db *sql.DB) *echo.Echo {
	go cleanDatabase(db)

	e := echo.New()
	e.AutoTLSManager.Prompt = autocert.AcceptTOS
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("stage.fomo.exchange")
	e.AutoTLSManager.Cache = autocert.DirCache("/mnt/fomo/autocert")

	middlewares.SetMainMiddlewares(e)

	service := k8s.NewService(
		micro.Name("api"))

	service.Init()

	// controllers
	authController := controllers.NewAuthController(db, service)
	sessionController := controllers.NewSessionController(db, service)
	userController := controllers.NewUserController(db, service)

	// required for health checks
	e.GET("/index.html", health)
	e.GET("/", health)

	// api group
	openApi := e.Group("/api")

	// open endpoints here
	openApi.POST("/login", authController.HandleLogin)
	openApi.POST("/signup", authController.HandleSignup)

	protectedApi := e.Group("/api")
	// set the auth middlewares
	protectedApi.Use(authController.RefreshAccess)
	middlewares.SetApiMiddlewares(protectedApi)

	// ###########################  protected endpoints here
	protectedApi.GET("/session", sessionController.HandleSession)
	protectedApi.GET("/logout", authController.HandleLogout)

	// user manangement endpoints
	protectedApi.PUT("/users/:userID/changepassword", userController.HandleChangePassword)
	protectedApi.PUT("/users/:userID", userController.HandleUpdateUser)

	go func() {
		if err := service.Run(); err != nil {
			log.Println("nope! ", err)
		}
	}()

	return e
}

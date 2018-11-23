package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/asciiu/oldiez/socket/controllers"
	"github.com/asciiu/oldiez/socket/middlewares"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

// send 200 ok to ping requests
func health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func NewRouter(db *sql.DB) *echo.Echo {
	e := echo.New()
	//e.AutoTLSManager.Prompt = autocert.AcceptTOS
	//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("stage.fomo.exchange")
	//e.AutoTLSManager.Cache = autocert.DirCache("/mnt/fomo/autocert")

	middlewares.SetMainMiddlewares(e)
	service := k8s.NewService(micro.Name("socket-service"))
	service.Init()

	// controllers
	socketController := controllers.NewWebsocketController()

	// websocket ticker
	e.GET("/ws", socketController.Connect)

	// required for health checks
	e.GET("/index.html", health)
	e.GET("/", health)

	go func() {
		if err := service.Run(); err != nil {
			log.Println("nope! ", err)
		}
	}()

	return e
}

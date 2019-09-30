package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/socket/controllers"
	"github.com/asciiu/appa/socket/middlewares"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	micro "github.com/micro/go-micro"
)

func health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func main() {
	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))

	if database, err := db.NewDB(dbURL); err != nil {
		log.Printf("ERROR: %s", err)
	} else {
		defer database.Close()

		echo := echo.New()
		//echo.AutoTLSManager.Prompt = autocert.AcceptTOS
		//echo.AutoTLSManager.HostPolicy = autocert.HostWhitelist("stage.fomo.exchange")
		//echo.AutoTLSManager.Cache = autocert.DirCache("/mnt/fomo/autocert")

		middlewares.SetMainMiddlewares(echo)

		// handle socket connections from socketController
		socketController := controllers.NewWebsocketController()

		// websocket ticker
		echo.GET("/ws", socketController.Connect)

		// required for health checks
		echo.GET("/index.html", health)
		echo.GET("/", health)

		service := k8s.NewService(micro.Name("socket-service"))
		service.Init()
		go func() {
			if err := service.Run(); err != nil {
				log.Println("nope! ", err)
			}
		}()

		// HTTPs server
		//echo.Logger.Fatal(echo.StartAutoTLS(":443"))
		port := fmt.Sprintf("%s", os.Getenv("SOCKET_PORT"))
		echo.Logger.Fatal(echo.Start(":" + port))
	}
}

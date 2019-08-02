// FOMO API
//
// Endpoints labeled open do not require authentication. The protected endpoints on the other hand, do require
// authentication. I'm not sure how long the jwt token should last. I'm thinking we should set the expire on
// that token to be super short - like 5 minutes (upto an hour maybe?) to minimize the amount of time an
// attacker can use that token. The refresh token will last longer - currently 7 days. If you make a request
// to a protected endpoint using a "Refresh" token in your request headers, you will receive a new
// authorization token (set-authorization) and refresh token (set-refresh) in the response headers when
// you make a request with an expired authorization token. You MUST replace both tokens in your request headers
// to stay authenticated. The old refresh token gets replaced on the backend therefore, you need to use the
// new refresh token to remain actively logged in.
//
//     Schemes: https
//     BasePath: /api
//     Version: 0.0.1
//     Author: The flo
//     Host: stage.fomo.exchange
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     Bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
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
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
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

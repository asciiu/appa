package middlewares

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func SetMainMiddlewares(e *echo.Echo) {

	e.Use(serverHeader)

	// this logs the server interaction
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method}  ${host}${path} ${latency_human}` + "\n",
	}))

	if isHTTPS, _ := strconv.ParseBool(os.Args[1]); isHTTPS {
		// use auto cert for ssl
		// note: this requires 'RUN apk --no-cache add ca-certificates' in the docker file
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(os.Getenv("DOMAIN_NAME"))
		e.AutoTLSManager.Cache = autocert.DirCache("/mnt/fluid/autocert")

		// redirect http to https
		e.Pre(middleware.HTTPSRedirect())
	}

	// TODO cors domain should read from config
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "Refresh"},
		ExposeHeaders: []string{"set-authorization", "set-refresh"},
	}))

	// don't crash on exceptions
	//e.Use(middleware.Recover())
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "plasma/0.0.1")
		return next(c)
	}
}

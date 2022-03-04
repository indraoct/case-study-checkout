package main

import (
	"case-study-checkout/handler/gql"
	"case-study-checkout/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
	"net/http"
	"os"
	"time"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	log.SetReportCaller(true)
}

func start(config config.Configuration, apiEcho *echo.Echo) {
	var err error

	log.Infof("--== Starting %s Server ==--", config.Name)

	apiEcho.Server.Addr = config.Port
	err = graceful.ListenAndServe(apiEcho.Server, 5*time.Second)

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}

func setWebRouter(config config.Configuration) *echo.Echo {
	e := echo.New()

	// Middleware Logger
	e.Use(middleware.Logger())
	root := e.Group(config.RootUrl)
	root.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG!")
	})

	root.Any("/graphql",func(c echo.Context) error {
		result := gql.GqlExecuteQuery(c)
		return c.JSON(http.StatusOK,&result)
	})

	return e
}

func configAndStartServer() {
	configMain := config.Get()

	//rest API router
	apiEcho := setWebRouter(configMain)
	start(configMain, apiEcho)
}

func main()  {
	initLogger()
	configAndStartServer()
}

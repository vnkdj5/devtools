package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vnkdj5/devtools/pkg/config"
	"github.com/vnkdj5/devtools/pkg/di"
	"github.com/vnkdj5/devtools/pkg/handlers"
	"github.com/vnkdj5/devtools/pkg/logging"
	"github.com/vnkdj5/devtools/pkg/router"
	"go.uber.org/zap"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error occucred while loading the config", err)
	}

	logger := logging.InitLogger()
	defer logger.Sync()
	logger.Sugar().Infof("Devtools Service. Environment: %s", config.Environment)
	// Echo instance
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Validator = conf.Validator

	container := di.BuildDIContainer(conf, logger)

	container.Provide(handlers.NewBase64Handler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `
       __                    __                       __       
  ____/ /  ___     _   __   / /_   ____     ____     / /  _____
 / __  /  / _ \   | | / /  / __/  / __ \   / __ \   / /  / ___/
/ /_/ /  /  __/   | |/ /  / /_   / /_/ /  / /_/ /  / /  (__  ) 
\__,_/   \___/    |___/   \__/   \____/   \____/  /_/  /____/  

-- Project by github.com/vnkdj5
`)
	})

	router.RegisterRoutes(e.Group("/api/v1"), container)

	// Start server
	go func() {
		if err := e.Start(":" + conf.HTTP.Port); err != nil && err != http.ErrServerClosed {
			logger.Fatal("shutting down the server. Reason: ", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/vnkdj5/devtools/pkg/config"
	"github.com/vnkdj5/devtools/pkg/handlers"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func RegisterRoutes(e *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		appConf *config.AppConfig,
		logger *zap.Logger,
	) {

		base64handler := handlers.NewBase64Handler(logger)
		base64Grp := e.Group("/base64")
		base64Grp.POST("/encode", base64handler.Encode)
		base64Grp.POST("/decode", base64handler.Decode)

	})

}

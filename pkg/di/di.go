package di

import (
	"github.com/vnkdj5/devtools/pkg/config"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func BuildDIContainer(conf *config.AppConfig, logger *zap.Logger) *dig.Container {
	container := dig.New()

	_ = container.Provide(func() *config.AppConfig {
		return conf
	})

	_ = container.Provide(func() *zap.Logger {
		return logger
	})

	return container
}

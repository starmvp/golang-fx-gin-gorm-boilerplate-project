package main

import (
	"boilerplate/pkg/loggers"
	"boilerplate/showcases/config/config"
	"boilerplate/showcases/config/db"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		loggers.Module,

		fx.Provide(zap.NewExample),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Invoke(
			func(db *gorm.DB, logger *zap.Logger) {
				logger.Debug("Database module invoked")
			},
			func(ddbc *config.DemoDBConfig, logger *zap.Logger) {
				logger.Debug("Config module invoked")

				fmt.Printf("ddbc: %#v\n", ddbc)
			},
			func(yc *config.YamlConfig, logger *zap.Logger) {
				logger.Debug("Config module invoked")

				fmt.Printf("yc: %#v\n", yc)
				fmt.Printf("ddbc: %#v\n", yc.GetSection("demo-db").(*config.DemoDBConfig))
			},
			func(logger *zap.Logger) {
				logger.Debug("Logger module invoked")
			},
		),
	)
	app.Run()
}

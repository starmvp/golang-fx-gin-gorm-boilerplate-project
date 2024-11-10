package main

import (
	"context"
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/utils"
	"golang-fx-gin-gorm-boilerplate-project/pkg/example"

	// "golang-fx-gin-gorm-boilerplate-project/pkg/example/server"

	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,

		logger.Module,

		example.Module,
		// server.Module,

		fx.Provide(zap.NewExample),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Invoke(
			func(db *gorm.DB, logger *zap.Logger) {
				logger.Debug("Database module invoked")
			},
			func(config *config.Config, logger *zap.Logger) {
				logger.Debug("Config module invoked")
			},
			func(logger *zap.Logger) {
				logger.Debug("Logger module invoked")
			},
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(fmt.Errorf("app.Start: %w", err))
	}

	/**
	 * Testing if the webserver is running
	 */
	go func() {
		time.Sleep(5 * time.Second)

		res, err := resty.
			New().
			R().
			Get(
				fmt.Sprintf("http://%s/ping", utils.GetWebserverAddr()),
			)
		if err != nil {
			log.Fatal(fmt.Errorf("resty.Get: %w", err))
		}
		fmt.Println("Testing Server: " + string(res.Body()))
	}()

	<-app.Wait()
}

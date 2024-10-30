package main

import (
	"context"
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/internal/app"
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"
	"golang-fx-gin-gorm-boilerplate-project/pkg/example"
	"golang-fx-gin-gorm-boilerplate-project/pkg/example/routers"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func getWebserverAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}
	// e.g 127.0.0.1:8080
	if !strings.Contains(port, ":") {
		port = ":" + port
	}
	return port
}

func main() {
	app := fx.New(
		app.Module,
		config.Module,
		db.Module,
		logger.Module,
		server.Module,

		example.Module,
		routers.Module,

		fx.Provide(zap.NewExample),
		fx.WithLogger(func(logger *logger.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Invoke(
			func(app *app.App, logger *logger.Logger) {
				logger.Debug("Webserver module invoked")
				go func() {
					_ = app.Server.Gin.Run(getWebserverAddr())
				}()
			},
			func(cl []*controller.Controller, logger *logger.Logger) {
				logger.Debug("Controller module invoked")
			},
			func(db *db.DB, logger *logger.Logger) {
				logger.Debug("Database module invoked")
			},
			func(config *config.Config, logger *logger.Logger) {
				logger.Debug("Config module invoked")
			},
			func(logger *logger.Logger) {
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
				fmt.Sprintf("http://%s/ping", getWebserverAddr()),
			)
		if err != nil {
			log.Fatal(fmt.Errorf("resty.Get: %w", err))
		}
		fmt.Println("Testing Server: " + string(res.Body()))
	}()

	<-app.Wait()
}

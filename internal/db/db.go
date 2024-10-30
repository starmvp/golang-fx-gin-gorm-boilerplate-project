package db

import (
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type DB struct {
	Logger   *logger.Logger
	DB       *gorm.DB
	DBConfig *gorm.Config
}

func New(
	Config *config.Config,
	Logger *logger.Logger,
	GormLogger logger.GormLogger,
) (*DB, error) {
	// TODO: add configure for db

	if Logger == nil {
		Logger = zap.NewNop()
	}

	fmt.Println("DB module invoked. GormLogger=", GormLogger)

	GormLogger.SetAsDefault()
	GormLogger.LogLevel = gormlogger.Warn

	db := DB{
		Logger: Logger,
		DBConfig: &gorm.Config{
			Logger: GormLogger,
		},
	}

	return &db, nil
}

var Module = fx.Provide(
	fx.Annotate(New),
)

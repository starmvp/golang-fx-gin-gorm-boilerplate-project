package db

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type DB struct {
	Logger   *logger.Logger
	DB       *gorm.DB
	DBConfig *gorm.Config
}

type DBParams struct {
	fx.In

	Config *config.Config
	Logger *logger.Logger
}

type DBResult struct {
	fx.Out

	DB *DB
}

func New(params DBParams) (DBResult, error) {
	// TODO: add configure for db

	l := params.Logger
	if l == nil {
		l = zap.NewNop()
	}

	gormLogger := zapgorm2.New(l)
	gormLogger.SetAsDefault()
	gormLogger.LogLevel = gormlogger.Warn

	db := DB{
		Logger: l,
		DBConfig: &gorm.Config{
			Logger: gormLogger,
		},
	}

	return DBResult{DB: &db}, nil
}

var Module = fx.Provide(New)

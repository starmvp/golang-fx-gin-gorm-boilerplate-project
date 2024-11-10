package db

import (
	"database/sql"
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func New(
	cfg *config.DBConfig,
	Logger *zap.Logger,
	GormLogger logger.GormLogger,
) *gorm.DB {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	sqlDB, err := sql.Open(cfg.Driver, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	GormLogger.SetAsDefault()
	GormLogger.LogLevel = gormlogger.Warn

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: GormLogger,
	})
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLife) * time.Second)

	return gormDB
}

var Module = fx.Provide(
	fx.Annotate(New),
)

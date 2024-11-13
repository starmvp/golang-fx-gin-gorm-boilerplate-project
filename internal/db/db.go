package db

import (
	"database/sql"
	"fmt"
	"time"

	"boilerplate/internal/config"
	"boilerplate/internal/logger"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(
	cfg *config.DBConfig,
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

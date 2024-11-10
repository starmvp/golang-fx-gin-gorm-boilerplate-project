package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthCheckService struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewHealthCheckService(db *gorm.DB, logger *zap.Logger) *HealthCheckService {
	return &HealthCheckService{
		DB:     db,
		Logger: logger,
	}
}

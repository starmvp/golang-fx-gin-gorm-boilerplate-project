package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PingService struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (s PingService) Name() string {
	return "ping-service"
}

func NewPingService(
	db *gorm.DB,
	logger *zap.Logger,
) *PingService {
	return &PingService{
		DB:     db,
		Logger: logger,
	}
}

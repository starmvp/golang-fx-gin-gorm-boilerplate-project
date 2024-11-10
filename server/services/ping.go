package services

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PingService struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewPingService(
	db *gorm.DB,
	logger *zap.Logger,
) *PingService {
	fmt.Println(">>> NewPingService")
	return &PingService{
		DB:     db,
		Logger: logger,
	}
}

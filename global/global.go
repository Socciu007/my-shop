package global

import (
	"my_shop/pkg/logger"
	"my_shop/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	GetDB *gorm.DB
)
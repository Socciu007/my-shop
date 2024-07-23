package initalize

import (
	"my_shop/global"
	"my_shop/pkg/logger"
)

// initializes the logger for the application
func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}

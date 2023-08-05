package global

import (
	"go-init/config"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	GO_CONFIG config.Configuration
	Go_VIPER  *viper.Viper
	Go_LOG    *zap.Logger
	GO_DB     *gorm.DB
)

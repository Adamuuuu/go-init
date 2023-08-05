package internal

import (
	"fmt"
	"go-init/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GO_CONFIG.App.DbType {
	case "mysql":
		logZap = global.GO_CONFIG.MySQL.LogZap
	}
	if logZap {
		global.Go_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

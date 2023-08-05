package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-init/core"
	"go-init/global"
	"go-init/initialize"
	"go.uber.org/zap"
)

const AppMode = "debug"

func main() {
	gin.SetMode(AppMode)

	// TODO：1.配置初始化
	global.Go_VIPER = core.InitializeViper()

	// TODO：2.日志
	global.Go_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.Go_LOG)

	global.Go_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))

	//  TODO：3.数据库连接
	global.GO_DB = initialize.Gorm()
	// TODO：4.其他初始化
	initialize.OtherInit()
	//启动服务
	fmt.Println("hello world")
	//r := gin.Default()
	//r.GET("", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})
	//address := fmt.Sprintf(":%d", global.EWA_CONFIG.App.Port)
	//r.Run(address)
	core.RunServer()
}

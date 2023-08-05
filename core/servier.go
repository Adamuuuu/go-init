package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-init/global"
	"go-init/initialize"
	"go.uber.org/zap"
	"net/http"
)

func RunServer() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	Router := initialize.Routers()
	r = Router
	address := fmt.Sprintf(":%d", global.GO_CONFIG.App.Port)
	global.Go_LOG.Info("server run success on ", zap.String("address", address))

	r.Run(address)

	// 保证文本顺序输出

}

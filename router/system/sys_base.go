package system

import (
	"go-init/model/system"
	"go-init/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")

	{
		//登录功能
		baseRouter.POST("login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
		//注册功能
		baseRouter.POST("register", func(context *gin.Context) {
			var form system.Register
			if err := context.ShouldBindJSON(&form); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": utils.GetErrorMsg(form, err),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	return baseRouter
}

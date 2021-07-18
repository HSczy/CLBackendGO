package route

import (
	"github.com/HSczy/CLBackendGO/controller"
	"github.com/gin-gonic/gin"
)

func BaseRoute(Route *gin.RouterGroup) (baseRoute gin.IRouter) {
	baseRoute = Route.Group("/")
	{
		baseRoute.POST("login", controller.LoginHandler)
		baseRoute.GET("captcha", controller.CaptchaHandler)
	}
	return
}

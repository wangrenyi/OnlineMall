package router

import (
	"github.com/gin-gonic/gin"
	"onlinemall/handler"
	"onlinemall/security"
)

func Router(router *gin.Engine) {

	routerGroup := router.Group("")
	{
		router.POST("/login", security.AuthLogin)
		router.POST("/register", security.RegisterUser)
	}
	v1 := routerGroup.Group("/v1")
	v1.Use(security.JWTMiddleware)

	initRouter(v1)
}

func initRouter(routerGroup *gin.RouterGroup) {
	//user handler
	handler.InitUserHandler(routerGroup)


}

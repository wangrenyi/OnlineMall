package router

import (
	"github.com/gin-gonic/gin"
	"onlinemall/db"
	"onlinemall/model"
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

	routerGroup.GET("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		user := new(model.MstUserInfo)
		connect := db.Connect()
		connect.First(user, id)
		context.JSON(200, gin.H{"data": &user})
	})
}

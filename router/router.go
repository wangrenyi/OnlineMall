package router

import (
	"github.com/gin-gonic/gin"
	"onlinemall/db"
	"onlinemall/model"
)

func Router(router *gin.Engine) {

	routerGroup := router.Group("/appbiz")

	routerGroup.GET("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		user := new(model.MstUserInfo)
		connect := db.Connect()
		connect.First(user, id)
		context.JSON(200, gin.H{"data": &user})
	})
}

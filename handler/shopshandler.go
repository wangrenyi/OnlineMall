package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
)

func InitShopsHandler(routerGroup *gin.RouterGroup) {
	shopsRouter := routerGroup.Group("/shops")
	{
		shopsRouter.GET("/list", getShopsByUserId)
		shopsRouter.GET("/detail/:shopsId", getShopsDetail)
		shopsRouter.POST("", saveShops)
		shopsRouter.DELETE("", deleteShops)
	}
}

func getShopsByUserId(context *gin.Context) {
	context.JSON(http.StatusOK, common.Info())
	return
}

func getShopsDetail(context *gin.Context) {
	context.JSON(http.StatusOK, common.Info())
	return
}

func saveShops(context *gin.Context) {
	context.JSON(http.StatusOK, common.Info())
	return
}

func deleteShops(context *gin.Context) {
	context.JSON(http.StatusOK, common.Info())
	return
}

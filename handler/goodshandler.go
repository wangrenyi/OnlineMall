package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
)

func InitGoodsHandler(routerGroup *gin.RouterGroup) {
	goodssRouter := routerGroup.Group("/goods")
	{
		goodssRouter.GET("/list", getGoodsByShops)
		goodssRouter.GET("/detail/:goodsId", getGoods)
		goodssRouter.POST("", saveGoods)
		goodssRouter.DELETE("", deleteGoods)
	}
}

func getGoodsByShops(context *gin.Context) {
	context.JSON(http.StatusOK, common.Ok())
	return
}

func getGoods(context *gin.Context) {
	context.JSON(http.StatusOK, common.Ok())
	return
}

func saveGoods(context *gin.Context) {
	context.JSON(http.StatusOK, common.Ok())
	return
}

func deleteGoods(context *gin.Context) {
	context.JSON(http.StatusOK, common.Ok())
	return
}

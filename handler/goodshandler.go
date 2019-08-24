package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
	"onlinemall/model"
	"onlinemall/repository"
)

func InitGoodsHandler(routerGroup *gin.RouterGroup) {
	goodssRouter := routerGroup.Group("/goods")
	{
		goodssRouter.GET("", getGoodsByShops)
		goodssRouter.GET("/detail/:goodsId", getGoods)
		goodssRouter.POST("", saveGoods)
		goodssRouter.DELETE("", deleteGoods)
	}
}

func getGoodsByShops(context *gin.Context) {

	var userId = context.Param("userId")

	shops := model.OnlineMallShops{}

	params := make(map[string]interface{}, 1)
	params["user_id"] = userId

	mstUserInfoDAO := repository.NewMstUserInfoDAO()
	mstUserInfoDAO.UniqueEntityByCondition(&shops, params)

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

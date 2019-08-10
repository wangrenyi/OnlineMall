package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
	"onlinemall/db"
	"onlinemall/model"
	"onlinemall/security"
	"time"
)

func InitUserHandler(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("/user")
	{
		userRouter.GET("/:loginid", getUserByLoginId)
		userRouter.POST("", updateUser)
		userRouter.DELETE("/:loginid", deleteUser)
	}
}

func getUserByLoginId(context *gin.Context) {
	loginid := context.Param("loginid")
	user := model.MstUserInfo{}
	connect := db.Connect()
	connect.Where("login_name = ?", loginid).First(&user)
	context.JSON(http.StatusOK, common.Success(user))
	return
}

func updateUser(context *gin.Context) {
	user := model.MstUserInfo{}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		context.Abort()
		return
	}
	loginUser := security.GetLoginUser()

	user.UpdateTime = time.Now()
	user.UpdateUser = loginUser.LoginName

	connect := db.Connect()
	connect.Model(&user).Where("login_name = ?", user.LoginName).Updates(user)
	context.JSON(http.StatusOK, common.Success(user))
	return
}

func deleteUser(context *gin.Context) {
	loginid := context.Param("loginid")

	user := model.MstUserInfo{}
	loginUser := security.GetLoginUser()
	user.UpdateUser = loginUser.LoginName
	user.UpdateTime = time.Now()
	user.LoginName = loginid
	user.Enabled = false

	connect := db.Connect()
	connect.Model(&user).Where("login_name = ?", user.LoginName).Updates(user)

	context.JSON(http.StatusOK, common.Success(user))
	return
}

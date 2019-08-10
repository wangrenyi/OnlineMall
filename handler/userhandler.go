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
		userRouter.GET("/:loginId", getUserByLoginId)
		userRouter.POST("", updateUser)
		userRouter.DELETE("/:loginId", deleteUser)
	}
}

func getUserByLoginId(context *gin.Context) {
	loginId := context.Param("loginId")
	user := model.MstUserInfo{}
	connect := db.Connect()
	connect.Where("login_name = ?", loginId).First(&user)
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
	tx := connect.Begin()
	if err := tx.Model(&user).Where("login_name = ?", user.LoginName).Updates(user).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	tx.Commit()

	context.JSON(http.StatusOK, common.Success(user))
	return
}

func deleteUser(context *gin.Context) {
	loginId := context.Param("loginId")

	user := model.MstUserInfo{}
	loginUser := security.GetLoginUser()
	user.UpdateUser = loginUser.LoginName
	user.UpdateTime = time.Now()
	user.LoginName = loginId
	user.Enabled = 0

	connect := db.Connect()
	tx := connect.Begin()
	if err := tx.Model(&user).Where("login_name = ?", user.LoginName).Updates(user).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	tx.Commit()

	context.JSON(http.StatusOK, common.Success(user))
	return
}

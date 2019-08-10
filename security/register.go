package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
	"onlinemall/db"
	"onlinemall/model"
	"time"
)

var loginUser = model.MstUserInfo{}

func AuthLogin(context *gin.Context) {
	defer common.Recover(context)

	user := model.MstUserInfo{}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	loginCheck(&user)

	token, err := GenerateToken(user)
	if err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	context.JSON(http.StatusOK, common.AuthSuccess(token))
	return
}

func RegisterUser(context *gin.Context) {
	defer common.Recover(context)

	registerUser := model.MstUserInfo{}
	if err := context.ShouldBindJSON(&registerUser); err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	existUser := model.MstUserInfo{}
	connect := db.Connect()
	connect.Where("login_name = ? and enabled = true ", loginUser.LoginName).First(&existUser)

	if existUser.LoginName == registerUser.LoginName && existUser.Password == registerUser.Password {
		common.PanicError("The user already exists.")
	} else {
		if loginUser.LoginName == "" {
			registerUser.CreateUser = registerUser.LoginName
			registerUser.UpdateUser = registerUser.LoginName
		} else {
			registerUser.CreateUser = loginUser.LoginName
			registerUser.UpdateUser = loginUser.LoginName
		}
		registerUser.CreateTime = time.Now()
		registerUser.UpdateTime = time.Now()
		registerUser.Enabled = true
		registerUser.Version = 1
		connect.Create(&registerUser)
	}

	context.JSON(http.StatusOK, common.Info())
	return
}

func loginCheck(user *model.MstUserInfo) {
	registerUser := model.MstUserInfo{}

	connect := db.Connect()
	connect.Where("login_name = ? and enabled = true ", user.LoginName).First(&registerUser)

	if registerUser.LoginName == "" {
		common.PanicError("The user does not exist!")
	}
	if registerUser.Password != user.Password || registerUser.LoginName != user.LoginName {
		common.PanicError("Incorrect account or password!")
	}

	loginUser = registerUser
}

func GetLoginUser() model.MstUserInfo {
	return loginUser
}

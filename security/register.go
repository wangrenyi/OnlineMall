package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
	"onlinemall/db"
	"onlinemall/model"
	"time"
)

var loginUser = new(model.MstUserInfo)

func AuthLogin(context *gin.Context) {
	defer common.Recover(context)

	if err := context.ShouldBindJSON(loginUser); err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	loginCheck(loginUser)

	token, err := GenerateToken(loginUser)
	if err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	context.JSON(http.StatusOK, common.AuthSuccess(token))
}

func RegisterUser(context *gin.Context) {
	defer common.Recover(context)

	registerUser := new(model.MstUserInfo)
	if err := context.ShouldBindJSON(registerUser); err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	existUser := new(model.MstUserInfo)
	connect := db.Connect()
	connect.Where("login_name = ? and enabled = true ", loginUser.LoginName).First(existUser)

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
		connect.Create(registerUser)
	}

	context.JSON(http.StatusOK, common.Info())
}

func loginCheck(loginUser *model.MstUserInfo) {
	user := new(model.MstUserInfo)

	connect := db.Connect()
	connect.Where("login_name = ? and enabled = true ", loginUser.LoginName).First(user)

	if user.LoginName == "" {
		common.PanicError("The user does not exist!")
	}
	if loginUser.Password != user.Password || loginUser.LoginName != user.LoginName {
		common.PanicError("Incorrect account or password!")
	}

	loginUser = user
}

func GetLoginUser() model.MstUserInfo {
	return *loginUser
}

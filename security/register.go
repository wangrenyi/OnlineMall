package security

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"onlinemall/common"
	"onlinemall/db"
	"onlinemall/model"
	"time"
)

var loginUser = model.MstUserInfo{}

func AuthLogin(context *gin.Context) {

	user := model.MstUserInfo{}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	err := loginCheck(&user)
	if err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	context.JSON(http.StatusOK, common.AuthSuccess(token))
	return
}

func RegisterUser(context *gin.Context) {

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
		registerUser.Enabled = 1
		registerUser.Version = 1

		tx := connect.Begin()
		if err := tx.Create(&registerUser).Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
			return
		}
		tx.Commit()
	}

	context.JSON(http.StatusOK, common.Ok())
	return
}

func loginCheck(user *model.MstUserInfo) error {
	registerUser := model.MstUserInfo{}

	connect := db.Connect()
	connect.Where("login_name = ? and enabled = true ", user.LoginName).First(&registerUser)

	if registerUser.LoginName == "" {
		return errors.New("The user does not exist!")
	}
	if registerUser.Password != user.Password || registerUser.LoginName != user.LoginName {
		return errors.New("Incorrect account or password!")
	}

	loginUser = registerUser
	return nil
}

func GetLoginUser() model.MstUserInfo {
	return loginUser
}

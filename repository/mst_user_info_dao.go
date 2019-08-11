package repository

import (
	"github.com/jinzhu/gorm"
	"onlinemall/db"
	"onlinemall/model"
)

type MstUserInfoDAO struct {
	Connect *gorm.DB
}

func NewMstUserInfoDAO() *MstUserInfoDAO {
	return &MstUserInfoDAO{db.Connect()}
}

func (mstUserInfoDAO *MstUserInfoDAO) SelectByLoginName(loginName string) *model.MstUserInfo {

	user := model.MstUserInfo{}
	connect := mstUserInfoDAO.Connect
	connect.Where("login_name = ?", loginName).First(&user)

	return &user
}

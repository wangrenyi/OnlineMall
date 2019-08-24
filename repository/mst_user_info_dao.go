package repository

import (
	"onlinemall/model"
)

type MstUserInfoDAO struct {
	BaseDAO
}

func NewMstUserInfoDAO() *MstUserInfoDAO {
	return &MstUserInfoDAO{NewBaseDAO()}
}

func (mstUserInfoDAO *MstUserInfoDAO) SelectByLoginName(loginName string) *model.MstUserInfo {

	user := model.MstUserInfo{}
	connect := mstUserInfoDAO.Connect
	connect.Where("login_name = ?", loginName).First(&user)

	return &user
}

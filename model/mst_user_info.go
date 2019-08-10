package model

import "time"

type MstUserInfo struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	LoginName   string    `gorm:"unique" json:"loginName"`
	Password    string    `json:"password"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Type        string    `json:"type"`
	Enabled     bool      `json:"enabled"`
	Remarks     string    `json:"remarks"`
	CreateUser  string    `json:"createUser"`
	CreateTime  time.Time `json:"createTime"`
	UpdateUser  string    `json:"updateUser"`
	UpdateTime  time.Time `json:"updateTime"`
	Version     uint      `json:"version"`
}

// 设置表名为`table`,不设置为tables
func (MstUserInfo) TableName() string {
	return "mst_user_info"
}

package model

import "time"

type OnlineMallShops struct {
	ID               string    `gorm:"primary_key" json:"id"`
	UserId           uint      `json:"userId"`
	ShopsName        string    `json:"shopsName"`
	ShopsRate        uint8     `json:"shopsRate"`
	BusinessCategory string    `json:"businessCategory"`
	Qualification    string    `json:"qualification"`
	Remarks          string    `json:"remarks"`
	CreateUser       string    `json:"createUser"`
	CreateTime       time.Time `json:"createTime"`
	UpdateUser       string    `json:"updateUser"`
	UpdateTime       time.Time `json:"updateTime"`
	Version          uint      `json:"version"`
}

// 设置表名为`table`,不设置为tables
func (OnlineMallShops) TableName() string {
	return "online_mall_shops"
}

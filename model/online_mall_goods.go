package model

import "time"

type OnlineMallGoods struct {
	ID         string    `json:"id"`
	ShopsId    string    `json:"shopsId"`
	GoodsName  string    `json:"goodsName"`
	Category   string    `json:"category"`
	Status     bool      `json:"status"`
	Remarks    string    `json:"remarks"`
	CreateUser string    `json:"createUser"`
	CreateTime time.Time `json:"createTime"`
	UpdateUser string    `json:"updateUser"`
	UpdateTime time.Time `json:"updateTime"`
	Version    uint      `version`
}

// 设置表名为`table`,不设置为tables
func (OnlineMallGoods) TableName() string {
	return "online_mall_goods"
}

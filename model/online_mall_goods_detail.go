package model

import "time"

//金额使用整型,避免浮点型计算不精确,以分为单位
type OnlineMallGoodsDetail struct {
	ID            uint      `json:"id"`
	GoodsId       string    `json:"goodsId"`
	origin        string    `json:"origin"`
	price         uint      `json:"price"`
	freight       uint      `json:"freight"`
	originalPrice uint      `json:"originalPrice"`
	specification string    `json:"specification"`
	remarks       string    `json:"remarks"`
	createUser    string    `json:"createUser"`
	createTime    time.Time `json:"createTime"`
	updateUser    string    `json:"updateUser"`
	updateTime    time.Time `json:"updateTime"`
	version       uint      `json:"version"`
}

// 设置表名为`table`,不设置为tables
func (OnlineMallGoodsDetail) TableName() string {
	return "online_mall_goods_detail"
}

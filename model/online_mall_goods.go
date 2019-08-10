package model

type OnlineMallGoods struct {
	ID        string `gorm:"primary_key" json:"id"`
	ShopsId   string `json:"shopsId"`
	GoodsName string `json:"goodsName"`
	Category  string `json:"category"`
	Status    bool   `json:"status"`
	Remarks   string `json:"remarks"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallGoods) TableName() string {
	return "online_mall_goods"
}

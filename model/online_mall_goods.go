package model

type OnlineMallGoods struct {
	ID        string `gorm:"primary_key" json:"id"`
	ShopsId   string `gorm:"unique_index" json:"shopsId"`
	GoodsName string `json:"goodsName"`
	Category  string `json:"category"`
	Status    uint8  `json:"status"`
	Remarks   string `json:"remarks"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallGoods) TableName() string {
	return "online_mall_goods"
}

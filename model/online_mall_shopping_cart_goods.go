package model

type OnlineMallShoppingCartGoods struct {
	ID             string `gorm:"primary_key" json:"id"`
	ShoppingCardId string `gorm:"index" json:"shoppingCardId"`
	GoodsId        string `json:"goodsId"`
	GoodsName      string `json:"goodsName"`
	Quantity       uint   `json:"quantity"`
	Status         uint8  `json:"status"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallShoppingCartGoods) TableName() string {
	return "online_mall_shopping_cart_goods"
}

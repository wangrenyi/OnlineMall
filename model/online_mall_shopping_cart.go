package model

type OnlineMallShoppingCart struct {
	ID     string `gorm:"primary_key" json:"id"`
	UserId uint   `gorm:"unique_index" json:"userId"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallShoppingCart) TableName() string {
	return "online_mall_shopping_cart"
}

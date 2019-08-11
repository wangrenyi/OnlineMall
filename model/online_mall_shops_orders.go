package model

type OnlineMallShopsOrders struct {
	ID           string `gorm:"primary_key" json:"id"`
	OrdersNumber string `gorm:"unique_index" json:"ordersNumber"`
	UserId       uint   `json:"userId"`
	AddressId    uint   `json:"addressId"`
	GoodsId      string `json:"goodsId"`
	GoodsName    string `json:"goodsName"`
	Quantity     string `json:"quantity"`
	PayWay       string `json:"payWay"`
	Remarks      string `json:"remarks"`
	Status       uint8  `json:"status"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallShopsOrders) TableName() string {
	return "online_mall_shops_orders"
}

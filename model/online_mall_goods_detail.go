package model

//金额使用整型,避免浮点型计算不精确,以分为单位
type OnlineMallGoodsDetail struct {
	ID            uint   `json:"id"`
	GoodsId       string `json:"goodsId"`
	Origin        string `json:"origin"`
	Price         uint   `json:"price"`
	Freight       uint   `json:"freight"`
	OriginalPrice uint   `json:"originalPrice"`
	Specification string `json:"specification"`
	Remarks       string `json:"remarks"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallGoodsDetail) TableName() string {
	return "online_mall_goods_detail"
}

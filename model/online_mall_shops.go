package model

type OnlineMallShops struct {
	ID               string `gorm:"primary_key" json:"id"`
	UserId           uint   `json:"userId"`
	ShopsName        string `json:"shopsName"`
	ShopsRate        uint8  `json:"shopsRate"`
	BusinessCategory string `json:"businessCategory"`
	Qualification    string `json:"qualification"`
	Status           uint8  `json:"status"`
	Remarks          string `json:"remarks"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallShops) TableName() string {
	return "online_mall_shops"
}

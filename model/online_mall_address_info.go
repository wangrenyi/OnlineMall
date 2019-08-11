package model

type OnlineMallAddressInfo struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	UserId        uint   `gorm:"unique_index" json:"userId"`
	ConsigneeId   uint   `json:"consigneeId"`
	ConsigneeName string `json:"consigneeName"`
	PhoneNumber   string `json:"phoneNumber"`
	Address       string `json:"address"`
	AddressDetail string `json:"addressDetail"`
	Default       uint8  `json:"default"`
	Status        uint8  `json:"status"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallAddressInfo) TableName() string {
	return "online_mall_address_info"
}

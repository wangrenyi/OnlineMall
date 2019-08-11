package model

type MstUserInfo struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	LoginName   string `gorm:"unique_index" json:"loginName"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Type        string `json:"type"`
	Enabled     uint8  `json:"enabled"`
	Remarks     string `json:"remarks"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (MstUserInfo) TableName() string {
	return "mst_user_info"
}

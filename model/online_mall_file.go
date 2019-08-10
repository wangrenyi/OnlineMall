package model

type OnlineMallFile struct {
	ID               string `gorm:"primary_key" json:"id"`
	ReqId            string `json:"reqId"`
	ReqType          string `json:"reqType"`
	FileOriginalName string `json:"fileOriginalName"`
	FileUniqueName   string `json:"fileUniqueName"`
	Subdirectory     string `json:"subdirectory"`
	PublicModel
}

// 设置表名为`table`,不设置为tables
func (OnlineMallFile) TableName() string {
	return "online_mall_file"
}

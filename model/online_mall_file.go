package model

import "time"

type OnlineMallFile struct {
	ID               string    `json:"id"`
	reqId            string    `json:"reqId"`
	reqType          string    `json:"reqType"`
	fileOriginalName string    `json:"fileOriginalName"`
	fileUniqueName   string    `json:"fileUniqueName"`
	subdirectory     string    `json:"subdirectory"`
	createUser       string    `json:"createUser"`
	createTime       time.Time `json:"createTime"`
	updateUser       string    `json:"updateUser"`
	updateTime       time.Time `json:"updateTime"`
	version          uint      `json:"version"`
}

// 设置表名为`table`,不设置为tables
func (OnlineMallFile) TableName() string {
	return "online_mall_file"
}

package model

import "time"

type PublicModel struct {
	CreateUser string    `json:"createUser"`
	CreateTime time.Time `json:"createTime"`
	UpdateUser string    `json:"updateUser"`
	UpdateTime time.Time `json:"updateTime"`
	Version    uint      `json:"version"`
}

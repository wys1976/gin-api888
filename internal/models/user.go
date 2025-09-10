package models

import (
	"time"
)

type User struct {
	ID         int64     `json:"userId"`
	Username   string    `json:"username"`
	Phone      string    `json:"phone"`
	Age        int       `json:"age"`
	CreateTime time.Time `json:"createTime"`
	Remark     string    `json:"remark"`
}

// UserRepository 用户数据访问接口[8](@ref)
type UserRepository interface {
	GetByID(userID int64) (*User, error)
	UpdatePhone(userID int64, phone string) error
}

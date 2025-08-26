package models

import (
	"database/sql"
	"time"
	"user_manager/database"
)

type User struct {
	ID         int64     `json:"userId"`
	Username   string    `json:"username"`
	Phone      string    `json:"phone"`
	Age        int       `json:"age"`
	CreateTime time.Time `json:"createTime"`
	Remark     string    `json:"remark"`
}

// GetUserByID 根据ID查询用户
func GetUserByID(userID int64) (*User, error) {
	query := `SELECT id, username, phone, age, create_time, remark FROM users WHERE id = ?`

	row := database.DB.QueryRow(query, userID)

	var user User
	var createTimeStr string

	err := row.Scan(&user.ID, &user.Username, &user.Phone, &user.Age, &createTimeStr, &user.Remark)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	// 解析时间字符串
	user.CreateTime, err = time.Parse("2006-01-02 15:04:05", createTimeStr)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserPhone 更新用户电话
func UpdateUserPhone(userID int64, phone string) error {
	query := `UPDATE users SET phone = ? WHERE id = ?`

	_, err := database.DB.Exec(query, phone, userID)
	if err != nil {
		return err
	}

	return nil
}

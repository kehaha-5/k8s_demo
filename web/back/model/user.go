package model

import "time"

type Gender string

var GenderMap = map[int64]Gender{
	1: "男",
	2: "女",
}

type User struct {
	ID        uint
	UserName  *string
	UserPhone uint
	Age       uint
	Sex       Gender
	CreatedAt *time.Time
	UpdatedAt *time.Time 
}

// 实现 Scan 接口，从数据库读取数据
func (g *Gender) Scan(value interface{}) error {
	*g = GenderMap[value.(int64)]
	return nil
}

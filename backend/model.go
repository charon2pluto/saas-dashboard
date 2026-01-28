package main

import (
	"time"

	"gorm.io/gorm"
)

// User 定义了用户数据模型 (符合 GORM 规范)
// 对应简历中的: "实现 GORM 模型映射"
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(100);unique_index" json:"username"`
	Password  string `gorm:"type:varchar(100)" json:"-"`
	Role      string `gorm:"default:'user'" json:"role"`
	LastLogin time.Time
}

// Stats 仪表盘数据模型
type Stats struct {
	ID        uint `gorm:"primaryKey"`
	UserCount int
	Revenue   float64
	Growth    string
}

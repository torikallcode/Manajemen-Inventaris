package models

import "time"

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Age      int
	CreateAt time.Time
	UpdateAt time.Time
}

package model

type Permission struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

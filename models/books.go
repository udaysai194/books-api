package models

import "gorm.io/gorm"

type Book struct {
	ID     uint     `gorm:"primary key;autoIncrement" json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}

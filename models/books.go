package models

import "gorm.io/gorm"

type Books struct {
	ID     uint     `json:"id"`
	Title  *string  `json:"title"`
	Author *string  `json:"author"`
	Price  *float64 `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}

package models

type Book struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// func MigrateBooks(db *pgxpool.Pool) error {
// 	err := db.AutoMigrate(&Book{})
// 	return err
// }

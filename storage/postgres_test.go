package storage

import (
	"books-api/models"
	"books-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresTest struct {
	DB     *pgxpool.Pool
	config *models.Config
}

func (pg PostgresTest) GetAllBooks(ctx *gin.Context) ([]models.Book, error) {
	books := []models.Book{}
	rows, err := pg.DB.Query(ctx, "SELECT * FROM books;")
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		utils.HandleError(err, "error fetching rows")
		books = append(books, book)
	}

	utils.HandleError(err, "no books found in database")
	return books, nil
}

func (pg PostgresTest) AddBooks(ctx *gin.Context, books []models.Book) error {

	for _, book := range books {
		pg.DB.Exec(ctx, "INSERT INTO books (title, author, price) VALUES ($1, $2, $3);", book.Title, book.Author, book.Price)
	}
	return nil
}

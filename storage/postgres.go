package storage

import (
	"books-api/models"
	"books-api/utils"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Template interface {
	GetAllBooks(ctx *gin.Context) ([]models.Book, error)
	AddBooks(ctx *gin.Context, books []models.Book) error
}

type Postgres struct {
	DB     *pgxpool.Pool
	config *models.Config
}

func ConfigPostgres(envFile string) (*models.Config, error) {
	err := godotenv.Load(envFile)
	utils.HandleError(err, "error loading .env file")

	config := &models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return config, err
}

func InitPostgres(config *models.Config) (Template, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.Connect(ctx, dsn)

	p := PostgresTest{
		DB:     db,
		config: config,
	}

	return p, err
}

func (pg Postgres) GetAllBooks(ctx *gin.Context) ([]models.Book, error) {
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

func (pg Postgres) AddBooks(ctx *gin.Context, books []models.Book) error {

	for _, book := range books {
		pg.DB.Exec(ctx, "INSERT INTO books (title, author, price) VALUES ($1, $2, $3);", book.Title, book.Author, book.Price)
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"books-api/models"
	"books-api/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Id     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type Repository struct {
	DB *gorm.DB
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(err)
	}
}

func main() {

	err := godotenv.Load(".env")
	handleError(err, "could not load .env file")

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.Connect(config)
	handleError(err, "could not connect to the database")

	err = models.MigrateBooks(db)
	handleError(err, "could not migrate")

	r := Repository{
		DB: db,
	}

	router := gin.Default()
	r.SetupRoutes(router)
	router.Run("localhost:8080")

}

func (r *Repository) SetupRoutes(router *gin.Engine) {
	router.POST("/create_books", r.AddBook)
	router.GET("/books", r.GetBooks)
}

func (r *Repository) GetBooks(c *gin.Context) {
	bookModels := &[]Book{}
	err := r.DB.Find(&bookModels)
	fmt.Println(bookModels)
	handleError(err.Error, "no books found in database")
	c.JSON(http.StatusOK, bookModels)
}

func (r *Repository) AddBook(c *gin.Context) {
	book := Book{}

	err := c.BindJSON(&book)
	handleError(err, "cant bind the books")
	err = r.DB.Create(&book).Error
	handleError(err, "cant POST books")
}

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

type Repository struct {
	DB *gorm.DB
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(err)
	}
}

func configure(envFile string) (*storage.Config) {
	err := godotenv.Load(envFile)
	handleError(err, "could not load .env file")

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return config
}

func main() {

	config := configure("mac.env")

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
	router.POST("/add-books", r.AddBooks)
	router.GET("/books", r.GetBooks)
	router.DELETE("/delete-book/:id", r.DeleteBookByID)
	router.GET("/book/:id", r.GetBookByID)
}

func (r *Repository) GetBooks(c *gin.Context) {
	books := &[]models.Book{}
	err := r.DB.Find(&books)
	handleError(err.Error, "no books found in database")
	c.JSON(http.StatusOK, books)
}

func (r *Repository) AddBooks(c *gin.Context) {
	books := &[]models.Book{}

	err := c.BindJSON(&books)
	handleError(err, "cant bind the books")
	err = r.DB.Create(&books).Error
	handleError(err, "cant POST books")
}

func (r *Repository) GetBookByID(c *gin.Context) {
	book := models.Book{}

	id := c.Param("id")
	err := r.DB.Where("id = ?", id).First(&book).Error
	handleError(err, "book with the given id not found")

	c.JSON(http.StatusOK, book)
}

func (r *Repository) DeleteBookByID(c *gin.Context) {
	book := models.Book{}

	id := c.Param("id")
	err := r.DB.Delete(book, id).Error
	handleError(err, "book with the given id not found")

	c.JSON(http.StatusOK, gin.H{"msg":"this worked"})
}
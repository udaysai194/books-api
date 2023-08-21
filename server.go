package main

import (
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

func Configure(envFile string) *storage.Config {
	err := godotenv.Load(envFile)
	HandleError(err, "could not load .env file")

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

func (r *Repository) SetupRoutes(router *gin.Engine) {
	router.POST("/add-books", r.AddBooks)
	router.GET("/books", r.GetBooks)
	router.DELETE("/delete-book/:id", r.DeleteBookByID)
	router.GET("/book/:id", r.GetBookByID)
}

func (r *Repository) GetBooks(c *gin.Context) {
	books := &[]models.Book{}
	err := r.DB.Find(&books)
	HandleError(err.Error, "no books found in database")
	c.JSON(http.StatusOK, books)
}

func (r *Repository) AddBooks(c *gin.Context) {
	books := &[]models.Book{}

	err := c.BindJSON(&books)
	HandleError(err, "cant bind the books")
	err = r.DB.Create(&books).Error
	HandleError(err, "cant POST books")
}

func (r *Repository) GetBookByID(c *gin.Context) {
	book := models.Book{}

	id := c.Param("id")
	err := r.DB.Where("id = ?", id).First(&book).Error
	HandleError(err, "book with the given id not found")

	c.JSON(http.StatusOK, book)
}

func (r *Repository) DeleteBookByID(c *gin.Context) {
	book := models.Book{}

	id := c.Param("id")
	err := r.DB.Delete(book, id).Error
	HandleError(err, "book with the given id not found")

	c.JSON(http.StatusOK, gin.H{"msg": "this worked"})
}

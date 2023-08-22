package main

import (
	"books-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) SetupRoutes(router *gin.Engine) {
	router.POST("/add-books", r.AddBooks)
	router.GET("/books", r.GetBooks)
	router.DELETE("/delete-book/:id", r.DeleteBookByID)
	router.GET("/book/:id", r.GetBookByID)
}

func (r *Repository) GetBooks(c *gin.Context) {
	books := []models.Book{}
	rows, err := r.DB.Query(c, "SELECT * FROM books;")

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		HandleError(err, "error fetching rows")
		books = append(books, book)
	}

	fmt.Println()
	HandleError(err, "no books found in database")
	c.JSON(http.StatusOK, books)
}

func (r *Repository) AddBooks(c *gin.Context) {
	// books := &[]models.Book{}

	// err := c.BindJSON(&books)
	// HandleError(err, "cant bind the books")
	// err = r.DB.Create(&books).Error
	// HandleError(err, "cant POST books")
}

func (r *Repository) GetBookByID(c *gin.Context) {
	// book := models.Book{}

	// id := c.Param("id")
	// err := r.DB.Where("id = ?", id).First(&book).Error
	// HandleError(err, "book with the given id not found")

	// c.JSON(http.StatusOK, book)
}

func (r *Repository) DeleteBookByID(c *gin.Context) {
	// book := models.Book{}

	// id := c.Param("id")
	// err := r.DB.Delete(book, id).Error
	// HandleError(err, "book with the given id not found")

	// c.JSON(http.StatusOK, gin.H{"msg": "this worked"})
}

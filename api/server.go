package api

import (
	"books-api/models"
	"books-api/storage"
	"books-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var database storage.Template

type Server struct {
	router *gin.Engine
	config *models.Config
}

func NewServer(config *models.Config) (*Server, error) {
	var srv Server
	var err error

	srv.config = config
	srv.router = gin.Default()
	database, err = storage.InitPostgres(config)
	utils.HandleError(err, "Erorr in connecting to postgress")

	return &srv, nil
}

func (s *Server) ListenAndServe(IP string, port string) {
	s.router.Run(IP + ":" + port)
}

func (s *Server) SetupRoutes() {
	s.router.POST("/add-books", s.AddBooks)
	s.router.GET("/books", s.GetBooks)
	s.router.DELETE("/delete-book/:id", s.DeleteBookByID)
	s.router.GET("/book/:id", s.GetBookByID)
}

func (s *Server) GetBooks(ctx *gin.Context) {
	books, err := database.GetAllBooks(ctx)
	utils.HandleError(err, "no books found in database")
	ctx.JSON(http.StatusOK, books)
}

func (r *Server) AddBooks(c *gin.Context) {
	// books := &[]models.Book{}

	// err := c.BindJSON(&books)
	// HandleError(err, "cant bind the books")
	// err = r.DB.Create(&books).Error
	// HandleError(err, "cant POST books")
}

func (r *Server) GetBookByID(c *gin.Context) {
	// book := models.Book{}

	// id := c.Param("id")
	// err := r.DB.Where("id = ?", id).First(&book).Error
	// HandleError(err, "book with the given id not found")

	// c.JSON(http.StatusOK, book)
}

func (r *Server) DeleteBookByID(c *gin.Context) {
	// book := models.Book{}

	// id := c.Param("id")
	// err := r.DB.Delete(book, id).Error
	// HandleError(err, "book with the given id not found")

	// c.JSON(http.StatusOK, gin.H{"msg": "this worked"})
}

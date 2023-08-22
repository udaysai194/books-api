package main

import (
	"books-api/storage"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(err)
	}
}

func main() {

	config := Configure("windows.env")

	db, err := storage.Connect(config)
	HandleError(err, "could not connect to the database")

	//err = models.MigrateBooks(db)
	//HandleError(err, "could not migrate")

	r := Repository{
		DB: db,
	}

	router := gin.Default()
	r.SetupRoutes(router)
	router.Run("localhost:8080")

}

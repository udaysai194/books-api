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

	var postgres storage.Postgres
	config, err := storage.ConfigPostgres("windows.env")
	HandleError(err, "errror is postgress configuration")
	postgres, err = storage.InitPostgres(config)
	HandleError(err, "Erorr in connecting to postgress")

	router := gin.Default()
	r.SetupRoutes(router)
	router.Run("localhost:8080")

}

package main

import (
	"books-api/api"
	"books-api/storage"
	"books-api/utils"
)

func main() {

	config, err := storage.ConfigPostgres("storage/windows.env")
	utils.HandleError(err, "errror in postgress configuration")

	server, err := api.NewServer(config)
	server.ListenAndServe("localhost", "8080")
	server.SetupRoutes()

}

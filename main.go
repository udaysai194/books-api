package main

import (
	"books-api/api"
	"books-api/utils"
)

func main() {

	server, err := api.NewServer()
	utils.HandleError(err, "error in starting new server")
	server.ListenAndServe("localhost", "8080")

}

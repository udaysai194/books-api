package handlingError

import (
	"fmt"
	"log"
)

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/fabiokusaba/go-rest-api-tutorial/cmd/api"
)

func main() {
	// initialize db
	// start api server
	apiServer := api.NewAPIServer(":8080")
	if err := apiServer.Run(); err != nil {
		log.Fatalln("error running api server")
	}
}

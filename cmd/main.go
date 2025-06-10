package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fabiokusaba/go-rest-api-tutorial/cmd/api"
	"github.com/fabiokusaba/go-rest-api-tutorial/config"
	"github.com/fabiokusaba/go-rest-api-tutorial/db"
)

func main() {
	// initialize db
	dbConn, err := db.NewPostgreSQLConnection(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.Host,
		config.Envs.Port,
		config.Envs.User,
		config.Envs.Password,
		config.Envs.DBName,
	))
	if err != nil {
		log.Fatalln("error connecting to postgres:", err)
	}

	if err := initDB(dbConn); err != nil {
		log.Fatalln("db connection failed:", err)
	}
	// start api server
	apiServer := api.NewAPIServer(":8080", dbConn)
	if err := apiServer.Run(); err != nil {
		log.Fatalln("error running api server")
	}
}

func initDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	log.Println("connected to database:", config.Envs.DBName)

	return nil
}

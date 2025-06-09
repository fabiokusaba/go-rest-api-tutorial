package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		"postgres://postgres:postgres123@localhost:5432/restapitutorial?sslmode=disable",
	)
	if err != nil {
		log.Fatalln("error initializing migration:", err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); err != nil && err == migrate.ErrNoChange {
			log.Println("migrate up error:", err)
		} else {
			log.Println("migrate up success")
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err == migrate.ErrNoChange {
			log.Println("migrate down error:", err)
		} else {
			log.Println("migrate down success")
		}
	}
}

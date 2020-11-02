package main

import (
	"log"

	"pismo/api"
	db "pismo/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	// dbSource      = "postgresql://pismo:password@localhost:5432/pismo_api?sslmode=disable"
	dbSource      = "postgresql://pismo:password@postgres:5432/pismo_api?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	store, err := db.NewStore(dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

package main

import (
	"lizee/pkg/products"
	"lizee/pkg/server"
	"lizee/pkg/storage"
)

const (
	serverPort = ":5000"

	host     = "172.28.1.1"
	dbport   = 5432
	user     = "dev"
	password = "Tx4BXPgfc]@;r"
	dbName   = "lizee"
)

func main() {
	// Init postgresql db
	postgres, err := storage.Connect(host, dbport, user, password, dbName)
	if err != nil {
		// Impossible to work without so we stop the program
		panic(err)
	}
	defer postgres.CloseConnection()

	// provide storage interface to products package
	products.InitAPIStorage(postgres)
	// Init Server
	serverInstance := server.Setup()
	serverInstance.Serve(serverPort)
}

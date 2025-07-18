package main

import (
	"database/sql"
	"log"

	"github.com/Amir228Kali/simplebank/api"
	db "github.com/Amir228Kali/simplebank/db/sqlc"
	"github.com/Amir228Kali/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize the database and server here
	// dbStore := db.NewStore(...)
	// server := api.NewServer(dbStore)
	// server.Start(":8080")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

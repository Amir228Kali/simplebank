package main

import (
	"database/sql"
	"log"

	"github.com/Amir228Kali/simplebank/api"
	db "github.com/Amir228Kali/simplebank/db/sqlc"
	"github.com/Amir228Kali/simplebank/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
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
	runDBMigration(config.MIGRATION_URL, config.DBSource)

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create migration:", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migration:", err)
	}
	log.Println("DB migration completed successfully")
}

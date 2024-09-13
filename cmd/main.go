package main

import (
	"database/sql"
	"fmt"
	"github.com/alfaysal/go-rest-api/cmd/api"
	"github.com/alfaysal/go-rest-api/config"
	db2 "github.com/alfaysal/go-rest-api/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	fmt.Println("calld from main")
	db, err := db2.NewMySQLStorage(mysql.Config{
		User:      config.Envs.DBUser,
		Passwd:    config.Envs.DBPassword,
		Addr:      config.Envs.DBAddress,
		DBName:    config.Envs.DBName,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8088", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")
}

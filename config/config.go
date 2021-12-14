package config

// TODO: uncomment below once you're ready to connect to Postgres

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	_ "github.com/lib/pq"
// )

var Port = 5050

// var SqlDB *sql.DB

// func ConnectToPostgres() {
// 	fmt.Println("Connecting to Postgres")

// 	const (
// 		host     = "localhost"
// 		port     = 5432
// 		user     = "root"
// 		password = "password"
// 		dbname   = "root"
// 	)

// 	dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", dbConnectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	SqlDB = db

// 	fmt.Println("Successfully connected to Postgres")
// }

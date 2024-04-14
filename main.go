package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	file, err := os.Open("setup.env")
	if err != nil {
		return
	}
	defer file.Close()
	 

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into key and value
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			continue // Skip invalid lines
		}
		key := pair[0]
		value := pair[1]

		// Set the environment variable
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get configuration values from environment variables
	dbServer := os.Getenv("DB_SERVER")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")

	
	server := dbServer
    port := dbPort
    user := dbUser
    password := dbPassword
    database := dbDatabase

    // Create a connection string
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
        server, user, password, port, database)

    // Create a new connection
	fmt.Println(connString)
    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Error creating connection pool: ", err.Error())
    }
    defer db.Close()

    // Use context to set timeout
    ctx := context.Background()
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: ", err.Error())
    }

    fmt.Println("Connected to the database!")


}
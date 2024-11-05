package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/loickcherimont/trucks/internal/config"
	"github.com/loickcherimont/trucks/internal/models"
	"github.com/loickcherimont/trucks/internal/utils"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to "dbname" database and execute a first query
func InitDB(dbName string, U *models.User) {
	// Retrieve sensitive data into .env file
	config.LoadVar(".env")
	var (
		err                 error
		query               string
		dbUsername          string = os.Getenv("DB_USERNAME")
		dbPassword          string = os.Getenv("DB_PASSWORD")
		adminLogin          string = os.Getenv("TRUCKS_USERNAME")
		adminPassword       string = os.Getenv("TRUCKS_PASSWORD")
		adminHashedPassword string = utils.HashPassword(adminPassword)
	)

	// ????
	dataSourceName := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbPassword, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Connection error 1: %q\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Connection error 2: %q\n", err)
	}

	log.Printf("Connection to DB %q with SUCCESS!\n", dbName)

	// Create a table and fill it with minimal content
	query = `
	CREATE TABLE IF NOT EXISTS user_admin(
		id INT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		PRIMARY KEY (id)
	);
	
	`
	_, err = db.Exec(query)

	if err != nil {
		log.Fatalf("Query error 1: %q\n", err)
	}
	fmt.Println("Query executed with SUCCESS!")

	// Fill it with user_admin login/password credentials
	query = `INSERT INTO user_admin(id, username, password) VALUES(?, ?, ?)`
	_, err = db.Exec(query, 1, adminLogin, adminHashedPassword)

	if err != nil {
		log.Printf("Query error on first INSERTION %q\n", err)
	}

	*U = models.User{
		Login:          adminLogin,
		Password:       adminPassword,
		HashedPassword: adminHashedPassword,
	}

	// Later :
	// Create a test to check if this user exists in database
	// Like check its "id" or other thing with SELECT sql statement
}

func FetchData(username string, password string, retrievedUser models.RetrievedUser) models.RetrievedUser {
	var db *sql.DB
	tableName := "user_admin"
	query := fmt.Sprintf(`SELECT id, username, password FROM %s WHERE (username = ?) AND (password = ?)`, tableName)
	if err := db.QueryRow(query, username, password).Scan(&retrievedUser.Login, &retrievedUser.Password); err != nil {
		log.Fatalf("Query error for FetchData %q", err)
	}
	return retrievedUser
}

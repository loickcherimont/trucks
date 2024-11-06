package database

import (
	"database/sql"
	"errors"
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
	defer db.Close()

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
		// Fix: Prevent duplication of the previous query
		// about user_admin table
		log.Printf("Query error on first INSERTION %q\n", err)
	}

	// Check if password and hash version are OK
	// Before store password and hash into U User struct
	if !utils.CheckHashPassword(adminHashedPassword, adminPassword) {
		log.Fatal(errors.New("hash and password are not the same"))
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

// func FetchData(username string, password string, retrievedUser models.RetrievedUser) models.RetrievedUser {
// 	var db *sql.DB
// 	tableName := "user_admin"
// 	query := fmt.Sprintf(`SELECT id, username, password FROM %s WHERE (username = ?) AND (password = ?)`, tableName)
// 	if err := db.QueryRow(query, username, password).Scan(&retrievedUser.Login, &retrievedUser.Password); err != nil {
// 		log.Fatalf("Query error for FetchData %q", err)
// 	}
// 	return retrievedUser
// }

func FetchAllData() []models.Truck {

	var (
		err error
		// query      string
		dbUsername string = os.Getenv("DB_USERNAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName     string = "db_transport"
		// adminLogin          string = os.Getenv("TRUCKS_USERNAME")
		// adminPassword string = os.Getenv("TRUCKS_PASSWORD")
		// adminHashedPassword string = utils.HashPassword(adminPassword)
		// db                  *sql.DB
	)

	// ????
	dataSourceName := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbPassword, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Connection error 1: %q\n", err)
	}
	defer db.Close()

	var (
		// err    error
		trucks []models.Truck
		// db     *sql.DB
	)

	rows, err := db.Query("SELECT * FROM trucks")
	if err != nil {
		log.Fatal("query error in database/database.go: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		var truck models.Truck
		err := rows.Scan(&truck.Id, &truck.FuelType, &truck.Payload, &truck.Distance)
		if err != nil {
			log.Fatal("scan error in database/database.go", err)
		}
		trucks = append(trucks, truck)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("rows error in database/database.go", err)
	}

	return trucks
}
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

func FetchAllData() []models.Truck {

	var (
		err        error
		dbUsername string = os.Getenv("DB_USERNAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName     string = "db_transport"
	)

	// ????
	dataSourceName := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbPassword, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Connection error 1: %q\n", err)
	}
	defer db.Close()

	var trucks []models.Truck

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

// Delete into DB a truck by its ID
func DeleteById(id int64) {

	var (
		err        error
		dbUsername string = os.Getenv("DB_USERNAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName     string = "db_transport"
	)

	// ????
	dataSourceName := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbPassword, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Connection error 1: %q\n", err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM trucks WHERE id = ?", id)
	if err != nil {
		log.Fatalln("Query error from database/database.go in DeleteById: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error in database/database.go at DeleteById func: ", err)
	}
	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d\n", rows)
	}
	fmt.Println("Truck deleted!")
}

// Add into DB the new truck using truck a model.Truck struct
// And update the array of trucks
func AddNewTruck(truck models.Truck) {
	var (
		err        error
		dbUsername string = os.Getenv("DB_USERNAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName     string = "db_transport"
	)

	// ????
	dataSourceName := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbPassword, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Connection error 1: %q\n", err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO trucks(id, fuel_type, payload, distance) VALUES(?, ?, ?, ?)", truck.Id, truck.FuelType, truck.Payload, truck.Distance)
	if err != nil {
		log.Fatalln("Query error from database/database.go in AddNewTruck: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error in database/database.go at AddNewTruck func: ", err)
	}
	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d\n", rows)
	}
	fmt.Println("New truck added!")
}

// Remarks:
// There are too much connection to DB
// Try to initialize once this operation!

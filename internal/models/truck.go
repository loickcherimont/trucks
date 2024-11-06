package models

import "github.com/gorilla/sessions"

// VARIABLES

// To initialize sessions
var (
	//????
	key   = []byte("secret")
	Store = sessions.NewCookieStore(key)

	// Trucks array for each Truck struct
	Trucks []Truck
)

// STRUCTS
type Truck struct {
	Id       int64
	FuelType string
	Payload  float64 // In tons
	Distance float64 // In kilometers
}

/*
	trucks := []models.Truck{
		{FuelType: "Diesel", Payload: 44, Distance: 500},
		{FuelType: "Gasoline", Payload: 19, Distance: 200},
		{FuelType: "Electricity", Payload: 3.5, Distance: 100},
	}
*/

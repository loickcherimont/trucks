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

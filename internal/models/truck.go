package models

import "github.com/gorilla/sessions"

// VARIABLES

// To initialize sessions
var (
	key   = []byte("secret")
	Store = sessions.NewCookieStore(key)
)

// STRUCTS
type Truck struct {
	FuelType string
	Payload  float64 // In tons
	Distance float64 // In kilometers
}

package models

type User struct {
	Login          string
	Password       string
	HashedPassword string
}

var U User

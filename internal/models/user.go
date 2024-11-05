package models

type User struct {
	Login          string
	Password       string
	HashedPassword string
}

type RetrievedUser struct {
	Login    string
	Password string
}

var U User

package utils

import (
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// UTILS
func ProcessError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Returns encoded password
// That is an hash of that encoded password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Password hashing error: %q", err)
	}

	return string(bytes)
}

// Returns a boolean indicating hash and password match
func CheckHashPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func ConvertStringIntoFloat(stringNumber string) float64 {
	floatNumber, err := strconv.ParseFloat(stringNumber, 64)
	if err != nil {
		log.Fatalf("Error conversion: %q", err)
	}
	return float64(floatNumber)
}

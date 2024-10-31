package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/loickcherimont/trucks/internal/models"
)

// To initialize sessions
var (
	key   = []byte("secret")
	store = sessions.NewCookieStore(key)
)

func main() {

	// Initialize new mux router
	router := mux.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Routes
	// Main ones
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/logout", logoutHandler)

	// Account routes
	router.HandleFunc("/admin", checkLogging(adminHandler))
	router.HandleFunc("/admin/trucks", checkLogging(trucksHandler))

	// Run the server
	fmt.Println("Server listening on: http://" + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

// HANDLERS

// Execute ./templates/index.html page without authentication
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, nil)
}

// GET: Execute ./templates/login.html page
// POST: Connect user to admin session
func loginHandler(w http.ResponseWriter, r *http.Request) {

	// To indicate user if he/she use wrong username/password
	var invalidCredentials bool

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == os.Getenv("TRUCKS_USERNAME") && password == os.Getenv("TRUCKS_PASSWORD") {
			session, err := store.Get(r, "cookie-name")
			processError(err, w)

			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		// User connects with wrong data
		invalidCredentials = true
	}
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	tmpl.Execute(w, struct{ InvalidCredentials bool }{invalidCredentials})
}

// Disconnect user from admin session
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	processError(err, w)

	session.Values["authenticated"] = false
	err = session.Save(r, w)
	processError(err, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Execute ./templates/admin.html page
func adminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/admin.html"))
	tmpl.Execute(w, nil)
}

// Execute ./templates/trucks.html page
func trucksHandler(w http.ResponseWriter, r *http.Request) {

	// Sample: data for trucks
	trucks := []models.Truck{
		{FuelType: "Diesel", Payload: 44, Distance: 500},
		{FuelType: "Gasoline", Payload: 19, Distance: 200},
		{FuelType: "Electricity", Payload: 3.5, Distance: 100},
	}

	tmpl := template.Must(template.ParseFiles("./templates/trucks.html"))
	tmpl.Execute(w, trucks)
}

// MIDDLEWARES

// Prevent user uses /admin/* without authentication
func checkLogging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "cookie-name")
		processError(err, w)

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}

// UTILS
func processError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

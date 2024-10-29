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
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/home", homeHandler)
	router.HandleFunc("/logout", logoutHandler)

	// Run the server
	fmt.Println("Server listening on: http://" + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

// HANDLERS
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD") {
			session, err := store.Get(r, "cookie-name")
			processError(err, w)

			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	tmpl.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	processError(err, w)

	session.Values["authenticated"] = false
	err = session.Save(r, w)
	processError(err, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "cookie-name")
	processError(err, w)

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("./templates/home.html"))
	tmpl.Execute(w, nil)
}

// UTILS
func processError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

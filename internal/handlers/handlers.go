package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/loickcherimont/trucks/internal/database"
	"github.com/loickcherimont/trucks/internal/models"
	"github.com/loickcherimont/trucks/internal/utils"
)

// VARIABLES
var (
	templatePath string = "./templates/*"
)

// Execute ./templates/index.html page without authentication
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(templatePath))
	utils.ProcessError(tmpl.ExecuteTemplate(w, "index.html", nil), w)
}

// GET: Execute ./templates/login.html page
// POST: Connect user to admin session
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var (
		// To indicate user if he/she uses wrong username/password
		invalidCredentials bool
	)

	if r.Method == http.MethodPost {

		// Login/password that user inputs
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Verify if user's inputs and user_admin data are the same
		if username == models.U.Login && password == models.U.Password {
			session, err := models.Store.Get(r, "session-name")
			utils.ProcessError(err, w)

			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		// User connects with wrong data
		invalidCredentials = true
	}
	tmpl := template.Must(template.ParseGlob(templatePath))
	utils.ProcessError(tmpl.ExecuteTemplate(w, "login.html", struct{ InvalidCredentials bool }{invalidCredentials}), w)
}

// Disconnect user from admin session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := models.Store.Get(r, "session-name")
	utils.ProcessError(err, w)

	session.Values["authenticated"] = false
	err = session.Save(r, w)
	utils.ProcessError(err, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Execute ./templates/admin.html page
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(templatePath))
	utils.ProcessError(tmpl.ExecuteTemplate(w, "admin.html", nil), w)
}

// Execute ./templates/trucks.html page
func TrucksHandler(w http.ResponseWriter, r *http.Request) {
	models.Trucks = database.FetchAllData()

	tmpl := template.Must(template.ParseGlob(templatePath))
	utils.ProcessError(tmpl.ExecuteTemplate(w, "trucks.html", models.Trucks), w)
}

// Delete a specific truck from DB
// Redirect user to /admin/trucks after deletion
func DeleteTruck(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalln("Conversion error on id: ", err)
	}
	database.DeleteById(int64(id), models.Trucks)

	http.Redirect(w, r, "/admin/trucks", http.StatusSeeOther)
}

// TODO - AddTruck func
func AddTruck(w http.ResponseWriter, r *http.Request) {

	// -- Instructions for future updates -- :

	// Fetch all submitted data + create new ID

	// Store data into a new {}Truck

	// Update DB (Add to DB the new truck)

	// Redirect user to /admin/trucks
}

// REMARKS :
// tmpl := template.Must(template.ParseGlob(templatePath))
// For each handlers all templates are parsed
// I think this is heavy and a bad practise!
// Better : parse all .tmpl files and one .html file
// Example : IndexHandler -> *.tmpl and index.html

package handlers

import (
	"html/template"
	"net/http"

	"github.com/loickcherimont/trucks/database"
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
		retrievedUser      models.RetrievedUser
	)

	if r.Method == http.MethodPost {

		// Login/password that user inputs
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if password and hash version are OK

		if !utils.CheckHashPassword(models.U.HashedPassword, password) {
			http.Error(w, "Hash and password wrong", http.StatusInternalServerError)
		}
		// Retrieve login/password from database
		retrievedUser = database.FetchData(username, password, retrievedUser)

		if username == retrievedUser.Login && password == retrievedUser.Password {
			session, err := models.Store.Get(r, "cookie-name")
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
	session, err := models.Store.Get(r, "cookie-name")
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

	// Sample: data for trucks
	trucks := []models.Truck{
		{FuelType: "Diesel", Payload: 44, Distance: 500},
		{FuelType: "Gasoline", Payload: 19, Distance: 200},
		{FuelType: "Electricity", Payload: 3.5, Distance: 100},
	}

	tmpl := template.Must(template.ParseGlob(templatePath))
	utils.ProcessError(tmpl.ExecuteTemplate(w, "trucks.html", trucks), w)
}

// REMARKS :
// tmpl := template.Must(template.ParseGlob(templatePath))
// For each handlers all templates are parsed
// I think this is heavy and a bad practise!
// Better : parse all .tmpl files and one .html file
// Example : IndexHandler -> *.tmpl and index.html

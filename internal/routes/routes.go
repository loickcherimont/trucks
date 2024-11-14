package routes

import (
	"github.com/gorilla/mux"
	"github.com/loickcherimont/trucks/internal/handlers"
	"github.com/loickcherimont/trucks/internal/middlewares"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()

	// Routes where authentication is not required
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/logout", handlers.LogoutHandler)

	// Routes with authentication
	// Controlled by middlewares
	router.HandleFunc("/admin", middlewares.CheckLogging(handlers.AdminHandler))
	router.HandleFunc("/admin/trucks", middlewares.CheckLogging(handlers.TrucksHandler))
	router.HandleFunc("/admin/trucks/delete/{id}", middlewares.CheckLogging(handlers.DeleteTruck))
	router.HandleFunc("/admin/trucks/add", middlewares.CheckLogging(handlers.AddTruck))

	return router
}

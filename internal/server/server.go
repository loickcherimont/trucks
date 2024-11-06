package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/loickcherimont/trucks/internal/database"
	"github.com/loickcherimont/trucks/internal/models"
	"github.com/loickcherimont/trucks/internal/routes"
)

func Run(addr string) {
	// Run the database before
	// tips: may be use goroutines?
	database.InitDB("db_transport", &models.U)
	router := routes.GetRoutes()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on: http://" + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"fmt"
	"github.com/GazzaMacD/workout_app/internal/app"
	"net/http"
	"time"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("We are running our app.")

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Health check route
	http.HandleFunc("/health", HealthCheck)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

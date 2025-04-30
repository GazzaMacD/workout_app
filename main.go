package main

import (
	"flag"
	"fmt"
	"github.com/GazzaMacD/workout_app/internal/app"
	"github.com/GazzaMacD/workout_app/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	// set port as flag option when running main so can change port
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running on %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

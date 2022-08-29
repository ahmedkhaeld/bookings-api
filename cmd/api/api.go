package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ahmedkhaeld/bookings-api/internals/configurations"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// App entry point to access App Configurations
var App configurations.Env

var env, _ = configurations.LoadConfig("../../.")

func main() {

	// start database
	db := run()
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {

		}
	}(db.SQL) // after main finishes shut down

	srv := &http.Server{
		Addr:              env.Port,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	App.InfoLog.Println(fmt.Sprintf("Starting back end server in %s mode on port %s", env.ENV, env.Port))

	//start the server
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}

	}()

	// gracefully shutdown the server when receiving a kill or interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		return
	}

}

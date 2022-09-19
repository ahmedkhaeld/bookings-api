package main

import (
	"github.com/ahmedkhaeld/bookings-api/internals/driver"
	handlers "github.com/ahmedkhaeld/bookings-api/internals/handlers/mongo"
	"log"
	"os"
)

var infoLog *log.Logger
var errorLog *log.Logger

func run() *driver.DB {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	App.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App.ErrorLog = errorLog

	// connect to database
	log.Println("Connecting to database")
	db := driver.ConnectMongo(env.MongoURI)

	// make sure the database connection is available to the handlers
	hand := handlers.Repo(db)
	handlers.New(hand)

	return db

}

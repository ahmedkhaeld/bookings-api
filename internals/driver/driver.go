package driver

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	ctx    context.Context
	err    error
	client *mongo.Client
)

//DB holds the database connection pool
type DB struct {
	Mongo *mongo.Client
}

//Conn a reference to the DB type
var Conn = &DB{}

func ConnectMongo(dsn string) *DB {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Conn.Mongo = client
	log.Println("Connected to MongoDB")
	return Conn

}

package driver

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// define some constraints to postgres connection
const (
	maxOpenDbConn = 10 // maximum open connections at any given time
	maxDbLifetime = 5 * time.Minute
	maxIdleDbConn = 5 // how many connection can remain in the pool but idle
)

var (
	ctx    context.Context
	err    error
	client *mongo.Client
)

//DB holds the database connection pool
type DB struct {
	Mongo *mongo.Client
	SQL   *sql.DB
}

//Conn a reference to the DB type
var Conn = &DB{}

// ConnectPostgres creates database pool for Postgres
func ConnectPostgres(dsn string) *DB {
	d, err := openPostgres(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	Conn.SQL = d

	err = pingDB(d)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...", err)
	}
	log.Println("Connected to Postgres")
	return Conn

}

// pingDB tries to ping the database
func pingDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// OpenPostgres creates a new database for the application
func openPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

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

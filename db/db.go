package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
	var err error
	connectionString := "host=localhost user=postgres password=q123 dbname=postgres sslmode=disable"
	Pool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalln("b5", err)
	}
}

package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
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

func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// Add hook to set session variables for every connection
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, `
			SET myapp.session.user_id = '';
			SET myapp.session.tenant_id = '';
			SET myapp.session.user_role = '';`)
		return err
	}

	return pgxpool.NewWithConfig(ctx, config)
}

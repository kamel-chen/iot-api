package database

import (
	"context"
	"log"
	"os"

	// _ "github.com/lib/pq"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPGDb()  (context.Context, *pgxpool.Pool) {
	host := os.Getenv("PG_HOST")
	password := os.Getenv("PG_PASSWORD")
	user := os.Getenv("PG_USER")
	name := os.Getenv("PG_DB_NAME")

	connStr := 
		"user=" + user + 
		" password=" + password + 
		" dbname=" + name + 
		" host=" + host +
		" sslmode=disable"
		
	ctx := context.Background()
	// db, err := sql.Open("postgres", connStr)
	db, err := pgxpool.Connect(ctx, connStr)

	if err != nil {
		log.Fatal(err)
	}

	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(980)

	return ctx, db
}

package db

import (
	"context"
	"fmt"
	"log"

	//_ "github.com/lib/pq"

	"github.com/jackc/pgx/v5/pgxpool"
	//_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	//HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Pool *pgxpool.Pool
}

func Initialize(username, password, host, database string) (Database, error) {
	/*
		db := Database{}
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			username, password, HOST, PORT, database)
		conn, err := sql.Open("pgx", dsn)
		if err != nil {
			return db, err
		}

		db.Conn = conn
		err = db.Conn.Ping()
		if err != nil {
			return db, err
		}
		log.Println("Database connection established")
		return db, nil
	*/

	db := Database{}
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, host, PORT, database)

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return db, err
	}

	db.Pool = pool
	err = db.Pool.Ping(context.Background())
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil

}

package database

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Conn *sql.DB
}

func ConnectPostgres() (*PostgresDB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")

	return &PostgresDB{
		Conn: db,
	}, nil
}

func (db *PostgresDB) Close() error {
	err := db.Conn.Close()
	if err != nil {
		return err
	}

	log.Println("Database connection closed")

	return nil
}

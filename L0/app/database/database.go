package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	user     string
	password string
	dbname   string
	host     string
	port     int
	conn     *sql.DB
}

const driverName = "postgres"

func New(user, password, dbname, host string, port int) *Database {

	db := Database{
		user:     user,
		password: password,
		dbname:   dbname,
		host:     host,
		port:     port,
	}

	return &db
}

// func (database *Database) Open(user, password, dbname, host string, port int) (*sql.DB, error) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := sql.Open(driverName, psqlInfo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	database.conn = db

//		return db, nil
//	}
func (database *Database) Open() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		database.host, database.port, database.user, database.password, database.dbname)

	db, err := sql.Open(driverName, psqlInfo)
	if err != nil {
		return nil, err
	}

	database.conn = db

	return db, nil
}

func (database *Database) Close() {
	if database.conn != nil {
		if err := database.conn.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		}
	}
}

func (database *Database) CreateTableIfExist(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders_json (" +
		"id serial primary key," +
		"data jsonb" +
		");")

	if err != nil {
		log.Println(err)
	}
}

func (database *Database) Insert(db *sql.DB, data []byte) {
	_, err := db.Exec("INSERT INTO orders_json (data) VALUES ($1)", data)
	if err != nil {
		log.Printf("Error inserting data into PostgreSQL: %v", err)
	}
}

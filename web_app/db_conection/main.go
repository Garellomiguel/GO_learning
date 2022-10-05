package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbconn = 5
const maxDbLifetime = 5 * time.Minute

func main() {

}

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDataBase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxIdleConns(maxIdleDbconn)
	d.SetConnMaxLifetime(maxDbLifetime)
	d.SetMaxOpenConns(maxOpenDbConn)

	dbConn.SQL = d
	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil

}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewDataBase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Example() {
	// Using jackc/pgx package
	// Connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=db_name user=user_name password=password")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to DB: %v", err))
	}
	defer conn.Close()
	// Test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Can not ping DB")
	}
	log.Println("Pinged DB")
	// Insert a row
	query := "INSERT INTO db.table (id, name) values ($1, $2)"
	_, err = conn.Exec(query, "I9", "Jack")
	if err != nil {
		log.Fatal(err)
	}
	// Get rows from table
	rows, err := conn.Query("SELECT id, name FROM db.table")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var id string
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Id:", id, "name:", name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("error scaning rows", err)
	}

	// Get row by id
	query = "SEELCT id, name from db.name where id = $1"
	row := conn.QueryRow(query, "I1")
	err = row.Scan(&id, &name)
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gsvaldevieso/go-dream-architecture/repository"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL() (*MySQL, error) {
	var ds = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open(os.Getenv("MYSQL_DRIVER"), ds)
	if err != nil {
		return &MySQL{}, nil
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the database")

	return &MySQL{db: db}, nil
}

//Exec is execute query
func (m *MySQL) Exec(query string, args ...interface{}) error {
	_, err := m.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

//Query returns results of a Query method.
func (m *MySQL) Query(query string, args ...interface{}) (repository.Row, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	row := NewMySQLRow(rows)

	return row, nil
}

//MySQLRow
type MySQLRow struct {
	Rows *sql.Rows
}

func NewMySQLRow(rows *sql.Rows) MySQLRow {
	return MySQLRow{Rows: rows}
}

//Scan returns results of a Scan method.
func (r MySQLRow) Scan(dest ...interface{}) error {
	if err := r.Rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

//Close returns results of a Close method.
func (r MySQLRow) Close() error {
	return r.Rows.Close()
}

//Next returns results of a Next method.
func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}

//Err returns results of a Err method.
func (r MySQLRow) Err() error {
	return r.Rows.Err()
}

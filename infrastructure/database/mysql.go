package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL stores the database structure
type MySQL struct {
	db *sql.DB
}

// NewMySQL return MySQL with database connection
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

	log.Println("Successfully connected to the MySQL database")

	return &MySQL{db: db}, nil
}

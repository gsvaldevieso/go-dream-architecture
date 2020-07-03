package adapter

import (
	"database/sql"

	"github.com/gsvaldevieso/go-dream-architecture/repository"
)

//Relational stores the database structure
type Relational struct {
	db *sql.DB
}

// NewRelational stores the database structure
func NewRelational(db *sql.DB) *Relational {
	return &Relational{db: db}
}

// Exec is execute query
func (m *Relational) Exec(query string, args ...interface{}) error {
	_, err := m.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Query returns results of a Query method.
func (m *Relational) Query(query string, args ...interface{}) (repository.Row, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	row := NewRelationalRow(rows)

	return row, nil
}

// RelationalRow stores the rows structure
type RelationalRow struct {
	Rows *sql.Rows
}

// Return NewRelational with database rows
func NewRelationalRow(rows *sql.Rows) RelationalRow {
	return RelationalRow{Rows: rows}
}

// Scan returns results of a Scan method.
func (r RelationalRow) Scan(dest ...interface{}) error {
	if err := r.Rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

// Close returns results of a Close method.
func (r RelationalRow) Close() error {
	return r.Rows.Close()
}

// Next returns results of a Next method.
func (r RelationalRow) Next() bool {
	return r.Rows.Next()
}

// Err returns results of a Err method.
func (r RelationalRow) Err() error {
	return r.Rows.Err()
}

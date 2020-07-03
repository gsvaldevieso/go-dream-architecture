package adapter

import (
	"database/sql"

	"github.com/gsvaldevieso/go-dream-architecture/repository"
)

type relational struct {
	db *sql.DB
}

// NewRelational return relational with database connection
func NewRelational(db *sql.DB) *relational {
	return &relational{db: db}
}

// Exec is execute query
func (m *relational) Exec(query string, args ...interface{}) error {
	_, err := m.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Query returns results of a Query method.
func (m *relational) Query(query string, args ...interface{}) (repository.Row, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	row := newRelationalRow(rows)

	return row, nil
}

type relationalRow struct {
	rows *sql.Rows
}

func newRelationalRow(rows *sql.Rows) relationalRow {
	return relationalRow{rows: rows}
}

// Scan returns results of a Scan method.
func (r relationalRow) Scan(dest ...interface{}) error {
	if err := r.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

// Close returns results of a Close method.
func (r relationalRow) Close() error {
	return r.rows.Close()
}

// Next returns results of a Next method.
func (r relationalRow) Next() bool {
	return r.rows.Next()
}

// Err returns results of a Err method.
func (r relationalRow) Err() error {
	return r.rows.Err()
}

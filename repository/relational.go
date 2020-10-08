package repository

// Relational is an abstraction for relational database
type Relational interface {
	Exec(string, ...interface{}) error
	Query(string, ...interface{}) (Row, error)
}

// Row is an abstraction for relational database result rows
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}

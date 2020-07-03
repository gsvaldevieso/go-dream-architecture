package repository

//Relational
type Relational interface {
	Exec(string, ...interface{}) error
	Query(string, ...interface{}) (Row, error)
}

//Row
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}

package data

import (
	"database/sql"
	"fmt"

	// TODO: move this to a main package
	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
)

var (
	// pool is the pool opened for the database
	pool *sql.DB
)

// GetPool is a safer interface for accessing the Pool
func GetPool() *sql.DB {
	if pool == nil {
		panic("GetPool() was used before the data.pool was defined." +
			" In other words, we can't connect to the database!" +
			" If this is happening in a test, did you use SetupTestDB() and" +
			" maybe ConnectToTestDB() in your func init()?")
	}

	return pool
}

// NewDB opens a standard DB
func NewDB() (*sql.DB, error) {

	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "metapods"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	return sql.Open("postgres", psqlInfo)
}

// SetupTestDB is used to setup a transactional database.
// Use it inside of an `init` function in a test file.
func SetupTestDB() {
	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "metapods_test"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	// we register an sql driver named "txdb"
	txdb.Register("txdb", "postgres", psqlInfo)
}

// NewTestDB creates a new of the test database
func NewTestDB() (*sql.DB, error) {
	return sql.Open("txdb", "identifier")
}

// ConnectToTestDB creates a new test db pool and sets it to data.pool
// Call this if you're using data.pool somewhere inside a function and want your test
// to use our test db.
func ConnectToTestDB() (*sql.DB, error) {
	db, err := NewTestDB()
	if err != nil {
		return db, err
	}

	pool = db
	return db, nil
}

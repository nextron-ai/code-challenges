package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	dsn := "host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	return sqlx.Connect("postgres", dsn)
}

//
// func InsertExample(db *sqlx.DB, col1, col2 string) error {
// 	query := `INSERT INTO example_table (col1, col2) VALUES ($1, $2)`
// 	_, err := db.Exec(query, col1, col2)
// 	return err
// }

// func SelectExample(db *sqlx.DB) ([]Result, error) {
//  // A struct "Result" precisa ter as tags "db" em seus campos
// 	var results []Result
// 	err := db.Select(&results, "SELECT col1, col2 FROM example_table")
// 	return results, err
// }

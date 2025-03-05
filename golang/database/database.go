package database

import (
	"codechalllenge/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	dsn := "host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	return sqlx.Connect("postgres", dsn)
}

func InsertPowerPlant(db *sqlx.DB, pp models.PowerPlant) error {
	query := `INSERT INTO power_plants (col1, col2) VALUES ($1, $2)`
	_, err := db.Exec(query, col1, col2)
	return err
}

//
// func InsertExample(db *sqlx.DB, col1, col2 string) error {
// 	query := `INSERT INTO example_table (col1, col2) VALUES ($1, $2)`
// 	_, err := db.Exec(query, col1, col2)
// 	return err
// }

// func SelectExample(db *sqlx.DB) ([]Result, error) {
// 	type Result struct {
// 		Col1 string `db:"col1"`
// 		Col2 string `db:"col2"`
// 	}
// 	var results []Result
// 	err := db.Select(&results, "SELECT col1, col2 FROM example_table")
// 	return results, err
// }

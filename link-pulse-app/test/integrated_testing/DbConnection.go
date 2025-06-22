package integrated_testing

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DbConnection() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)

	log.Println("The DSN was : ", dsn)

	return sql.Open("postgres", dsn)
}

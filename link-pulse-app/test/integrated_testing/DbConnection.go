package integrated_testing

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func DbConnection() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	log.Printf("Connecting to database with user: %s and password: %s and dbname: %s", user, password, dbname)
	dsn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)

	log.Println("The DSN was : ", dsn)

	return sql.Open("postgres", dsn)
}

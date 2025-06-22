package integrated_testing

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DbConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println("The DSN was : ", dsn)

	return sql.Open("postgres", dsn)
}

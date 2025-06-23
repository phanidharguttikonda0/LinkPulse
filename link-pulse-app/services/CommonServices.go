package services

import (
	"database/sql"
	"log"
)

func CheckPremium(db *sql.DB, userId int, num string) (bool, error) {
	var premium string
	err := db.QueryRow("select premium from Users where id = $1", userId).Scan(&premium)
	if err != nil {
		log.Println("Error Occurred in CheckPremium Service was ", err)
		return false, err
	}

	if premium == "1" && num == "1" {
		return true, nil
	} else if premium == "2" && num == "2" {
		return true, nil
	} else if premium == "3" {
		return true, nil
	}

	return false, nil
}

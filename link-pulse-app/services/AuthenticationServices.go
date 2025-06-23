package services

import (
	"database/sql"
	"github.com/phanidharguttikonda0/LinkPulse/models"
	"log"
)

// NewUser * here 0 means free user , 1 means only for website urls, 2 for only file sharing, 3 for both for premium feild
func NewUser(db *sql.DB, signUp *models.NewUser) (bool, int) {
	log.Println("Inserting a new User")
	var id int
	query := `insert into Users (mail_id, mobile, username, password) values ($1, $2, $3, $4) returning id`
	err := db.QueryRow(query, signUp.MailId, signUp.Mobile, signUp.User.Username, signUp.User.Password).Scan(&id)
	if err != nil {
		log.Printf("Error while inserting user: %v", err)
		return false, -1
	}
	return true, id
}

func CheckUser(db *sql.DB, signIn *models.User) (bool, int) {
	log.Println("Checking User")
	query := `select id from Users where username = $1 and password = $2`
	row := db.QueryRow(query, signIn.Username, signIn.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return false, -1
	}
	return true, id
}

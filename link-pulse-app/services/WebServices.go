package services

import (
	"database/sql"
	"log"
)

func CustomNameCheckService(db *sql.DB, CustomName string) (bool, error) {
	var name string
	err := db.QueryRow("SELECT shorten_url from website_urls where shorten_url=$1", CustomName).Scan(&name)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func NewUrl(db *sql.DB, orginal string, name string, userId int) (string, error) {
	log.Println("Storing the New Url as Service")
	var shortenUrl string
	err := db.QueryRow("Insert into website_urls (user_id, original_url, shorten_url) VALUES ($1, $2, $3) RETURNING shorten_url",
		userId, orginal, name).Scan(&shortenUrl)

	if err != nil {
		log.Println("Error occurred while Storing New Url ", err)
		return "", err
	}
	log.Println("Successfully Executed the Insert Query for Website Urls")
	return shortenUrl, nil
}

func StoringPremiumUserInsights() {
	// here we are going to use the dynamo db
}

func GetOriginalUrl(db *sql.DB, ShortenUrl string) (string, error) {
	var originalURL string
	err := db.QueryRow(`
    UPDATE website_urls
    SET clicks = clicks + 1
    WHERE shorten_url = $1
    RETURNING original_url
`, ShortenUrl).Scan(&originalURL)

	if err != nil {
		log.Println("Error occurred while Getting OriginalUrl ", err)
		return "", err
	} else {
		return originalURL, nil
	}
}

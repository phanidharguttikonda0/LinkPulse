package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/services"
	"log"
	"net/http"
	"time"
)

func PostUrlShortner(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		url := c.Query("url")

		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url query parameter is required"})
			return
		}

		CustomName := c.MustGet("CustomName").(string)

		userId := c.MustGet("UserId").(int)

		ShortenUrl, err := services.NewUrl(db, url, CustomName, userId)
		if err != nil {
			log.Println(err, " This Error Occurred while Storing a new Shorten Url")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		log.Println("Successfully Stored the new Shorten Url ", ShortenUrl)
		c.JSON(http.StatusCreated, gin.H{"url": ShortenUrl})
	}
}

// UrlShortner In front-end we need attach domain/shorten_url
func UrlShortner(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		url := c.Query("url")

		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url query parameter is required"})
			return
		}

		userId := c.MustGet("UserId").(int)
		var finalHash string
		// here we are going to generate a unique name using the original url along with current time stamp an hash of it

		timestamp := time.Now().Nanosecond() // for better precision
		data := fmt.Sprintf("%s%d", url, timestamp)
		hash := sha256.Sum256([]byte(data))
		// Encoding to base64 then trim to 8-12 characters for URL
		shortHash := base64.URLEncoding.EncodeToString(hash[:])

		finalHash = shortHash[:12] // we are trimming to 12 characters
		// up to 92 billion hashes can be stored with out collision for 12 characters trimming

		log.Println("The Final Hash trimmed to 12 was ", finalHash)

		ShortenUrl, err := services.NewUrl(db, url, finalHash, userId)
		if err != nil {
			log.Println(err, " This Error Occurred while Storing a new Shorten Url")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		log.Println("Successfully Stored the new Shorten Url ", ShortenUrl)
		c.JSON(http.StatusCreated, gin.H{"url": ShortenUrl})
	}
}

func CheckCustomName(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		CustomName := c.Param("name")

		value, err := services.CustomNameCheckService(db, CustomName)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusOK, gin.H{"Exists": value})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"Exists": value})
		}
		/*
			The Front-End needs to allow only if the Exists was false, such that the Post API of the
			Url Short-ner Route will be Executed Automatically after entering Proceed. If Exists was False
			Then it should no Allow the Proceed to takes place.
		*/
	}
}

package integrated_testing

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestingwebsiteUrls(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, err := DbConnection()
	if err != nil {
		t.Errorf("Error connecting to db: %v", err)
		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Errorf("Error closing db: %v", err)
			return
		}
	}()

	// we don't know in which order does the tests execute so pushing to the database
	var userId int
	// Insert user required for the test
	err = db.QueryRow(`INSERT INTO Users (username, password, mail_id, mobile, premium) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		"phanianna", "phanianna", "phanianna@gmail.com", "+918885558760", "1").Scan(&userId)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
		return
	} // this user has the premium

	router := gin.Default()
	routes.WebRoutes(router, db, "phani")
	header, err := middlewares.CreateAuthorizationHeader("phani", userId, "phanianna")
	if err != nil {
		t.Errorf("Error creating header: %v", err)
		return
	}

	form := url.Values{}
	form.Add("Name", "phani")

	req1, _ := http.NewRequest("POST", "/website/url-shortner?url=https://github.com/phanidharguttikonda0/LinkPulse/actions", strings.NewReader(form.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req1.Header.Add("Authorization", header)

	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)

	if resp1.Code != http.StatusOK {
		t.Errorf("Expected status 200, for resp1 but got %d", resp1.Code)
		return
	}

	req2, _ := http.NewRequest("GET", "/website/", nil)
	req2.Header.Add("Authorization", header)

	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)

	if resp2.Code != http.StatusOK {
		t.Errorf("Expected status 200 for resp2 but, got the following %d", resp2.Code)
		return
	}

	var response1 map[string]interface{}
	err = json.Unmarshal(resp1.Body.Bytes(), &response1)
	if err != nil {
		t.Fatalf("Failed to parse JSON response for resp1: %v", err)
	}

	var response2 map[string]interface{}
	err = json.Unmarshal(resp2.Body.Bytes(), &response2)
	if err != nil {
		t.Fatalf("Failed to parse JSON response for resp2: %v", err)
	}

	val1, ok1 := response1["url"].(string)
	val2, ok2 := response2["url"].(string)

	if !ok1 {
		t.Errorf("Failed to parse JSON response for resp1: %v", response1)
		return
	}

	if !ok2 {
		t.Errorf("Failed to parse JSON response for resp2: %v", response2)
		return
	}

	log.Println("The response of the resp1 and resp2 are ", val1, val2)

}

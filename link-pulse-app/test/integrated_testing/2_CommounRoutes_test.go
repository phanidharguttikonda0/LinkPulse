package integrated_testing

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// jwt secret was "phani" in over all the tests
func TestIsPremiumRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, err := DbConnection()
	if err != nil {
		t.Fatalf("Error connecting to test database: %v", err)
	}
	defer db.Close()
	// we don't know in which order does the tests execute so pushing to the database
	var userId int
	// Insert user required for the test
	err = db.QueryRow(`INSERT INTO Users (username, password, mail_id, mobile, premium) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		"Phaniidhar", "Phaniidhar", "phanidharreddy@gmail.com", "+918885758760", "0").Scan(&userId)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
		return
	}

	router := gin.Default()
	routes.CommonRoutes(router, db, "phani")

	// Generate valid token
	header, err := middlewares.CreateAuthorizationHeader("phani", userId, "Phaniidhar")
	if err != nil {
		t.Fatalf("Failed to create Authorization header: %v", err)
	}

	req, _ := http.NewRequest("GET", "/common/is-premium/1", nil)
	req.Header.Add("Authorization", header)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	contentType := resp.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	isPremium, ok := response["isPremium"].(bool)
	if !ok || !isPremium {
		t.Errorf("Expected isPremium=true, got: %v", response["isPremium"])
	}

	// Clean up DB
	_, _ = db.Exec(`DELETE FROM users WHERE id = $1`, 1)
}

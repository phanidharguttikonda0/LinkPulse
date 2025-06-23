package integrated_testing

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// jwt secret was "phani" in over all the tests
func TestIsPremiumRoute(t *testing.T) {
	// 2 routes one
	gin.SetMode(gin.TestMode)

	db, err := DbConnection()
	if err != nil {
		t.Fatalf("Error connecting to test database: %v", err)
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			t.Errorf("Error closing connection: %v", err)
		}
	}(db)

	router := gin.Default()
	routes.CommonRoutes(router, db, "phani")

	req, _ := http.NewRequest("GET", "/common/is-premium/1", nil)
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTA3Njc0NTksInVzZXJJZCI6MSwidXNlcm5hbWUiOiJQaGFuaWRoYXIifQ.yze0JE9UQvYo69G7It3rTNSJ7o5BAcgxAVjUzZ59X7s") // it was a created authorized key
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("SignUpHandler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
		return
	}
	log.Println("The response was ")
	log.Println(resp.Result())
	log.Println("Let's Check the Result ")

	// Header Check
	contentType := resp.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
		return
	}

	// Body JSON Check
	var response map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	isPremium, ok := response["isPremium"].(bool)
	if !ok || !isPremium {
		t.Errorf("Expected isPremium=true, got: %v", response["isPremium"])
	}

}

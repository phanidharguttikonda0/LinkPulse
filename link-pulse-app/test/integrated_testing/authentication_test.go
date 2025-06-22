package integrated_testing

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSignUpHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log.Println("Test Mode was Set in SignUpHandler")

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
	routes.AuthenticationRoutes(router, db, "phani")

	form := url.Values{}
	form.Set("user.username", "Phanidhar") // because of data bindings we need to user property name first
	form.Set("user.password", "Phanidhar")
	form.Set("mailId", "phanidhar@gmail.com")
	form.Set("mobile", "+918885858760")

	req, _ := http.NewRequest("POST", "/authentication/sign-up", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("SignUpHandler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
	}

}

func TestSignInHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log.Println("Test Mode was Set in SignInHandler")
	db, err := DbConnection()
	if err != nil {
		t.Fatalf("Error connecting to test database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	router := gin.Default()
	log.Println("Specified Route by passing db and jwt secret")
	routes.AuthenticationRoutes(router, db, "phani")
	/*
		Here we have set-up the actual routes, by calling the authentication routes
	*/

	form := url.Values{}
	form.Set("username", "Phanidhar")
	form.Set("password", "Phanidhar")

	req, _ := http.NewRequest("POST", "/authentication/sign-in", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	/*
		📌 Sends the fake request to your Gin router and records the response in resp (acts like an in-memory HTTP server).
	*/

	if resp.Code != http.StatusOK {
		t.Errorf("SignInHandler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
	}

}

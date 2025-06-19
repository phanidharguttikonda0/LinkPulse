package unit_testing

import (
	"fmt"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"testing"
)

func TestAuthorizationCheck(t *testing.T) {
	token, err := middlewares.CreateAuthorizationHeader(1, "phani")
	if err != nil {
		t.Errorf("Error creating token: %v", err)
	}
	claims, err := middlewares.AuthorizationCheck(token)
	if err != nil {
		t.Errorf("Error verifying token: %v", err)
	}
	fmt.Println(claims)
}

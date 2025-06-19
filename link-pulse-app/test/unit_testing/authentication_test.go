package unit_testing

import (
	"errors"
	"fmt"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"regexp"
	"testing"
)

func TestAuthorizationCheck(t *testing.T) {
	token, err := middlewares.CreateAuthorizationHeader("hehe", 1, "phani")
	if err != nil {
		t.Errorf("Error creating token: %v", err)
	}
	claims, err := middlewares.AuthorizationCheck("hehe", token)
	if err != nil {
		t.Errorf("Error verifying token: %v", err)
	}
	fmt.Println(claims)
}

// Define your regex patterns (example patterns, adjust as needed)
var (
	UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)                             // Example: 3-16 alphanumeric or underscore
	MobileRegex   = regexp.MustCompile(`^[0-9]{10}$`)                                      // Example: 10 digits
	EmailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`) // Standard email
)

// User struct for sign-in
type User struct {
	Username string
	Password string
}

// NewUser struct for sign-up
type NewUser struct {
	Username string
	Password string
	Mobile   string
	MailId   string
}

// isValid checks if a string matches a given regex
func isValid(regex *regexp.Regexp, s string) bool {
	return regex.MatchString(s)
}

// SignInValidation validates user sign-in credentials
func (signIn *User) TestSignInValidation() (bool, error) {
	if isValid(UsernameRegex, signIn.Username) {
		if len(signIn.Password) >= 8 {
			return true, nil
		} else {
			return false, errors.New("password should be at least 8 characters")
		}
	} else {
		return false, errors.New("username is invalid")
	}
}

// SignUpValidation validates new user sign-up credentials
func (signUp *NewUser) SignUpValidation() (bool, error) {
	if isValid(UsernameRegex, signUp.Username) {
		if len(signUp.Password) >= 8 {
			if isValid(MobileRegex, signUp.Mobile) {
				if isValid(EmailRegex, signUp.MailId) {
					return true, nil
				} else {
					return false, errors.New("invalid mail id")
				}
			} else {
				return false, errors.New("invalid mobile number")
			}
		} else {
			return false, errors.New("password should be at least 8 characters")
		}
	} else {
		return false, errors.New("username is invalid")
	}
}

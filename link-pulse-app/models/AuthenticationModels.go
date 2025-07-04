package models

import (
	"errors"
	"log"
	"regexp"
)

type NewUser struct {
	User   User   `form:"user"`
	MailId string `form:"mailId"`
	Mobile string `form:"mobile"`
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// now we need to username or Username as we mentioned in struct tags
// Starting with small case means private to the package , starting with Capital case are public and can be used
// in any package so Username where starting was Capital so i can use it any package

const EmailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
const MobileRegex = `^\+\d{1,3}\d{10}$`
const UsernameRegex = `^[a-zA-Z_][a-zA-Z0-9_]{2,}$`

func IsValid(pattern string, value string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func (signIn *User) SignInValidation() (bool, error) {
	log.Printf("username was %s and password was %s", signIn.Username, signIn.Password)
	if IsValid(UsernameRegex, signIn.Username) {
		if len(signIn.Password) >= 8 {
			return true, nil
		} else {
			return false, errors.New("password should be at least 8 characters")
		}
	} else {
		return false, errors.New("username is invalid")
	}
}

func (signUp *NewUser) SignUpValidation() (bool, error) {
	log.Printf("Mobile was %s and MailId was %s", signUp.Mobile, signUp.MailId)
	log.Println("The Sign Up was ", signUp)
	log.Printf("Username was %s and the Password was %s ", signUp.User.Username, signUp.User.Password)
	if IsValid(MobileRegex, signUp.Mobile) {
		// let's check is country id is valid or not
		value := len(signUp.Mobile) - 10
		CountryCode := signUp.Mobile[0:value]
		log.Println("Country Code was", CountryCode)
		if IsCountryCodeCorrect(CountryCode) {
			if IsValid(EmailRegex, signUp.MailId) {
				_, err := signUp.User.SignInValidation()
				if err != nil {
					return false, err
				}
				return true, nil
			} else {
				return false, errors.New("invalid mail id")
			}
		} else {
			return false, errors.New("Invalid Country Code")
		}

	} else {
		return false, errors.New("invalid mobile number")
	}
}

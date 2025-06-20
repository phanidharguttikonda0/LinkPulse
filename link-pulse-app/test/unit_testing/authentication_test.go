package unit_testing

import (
	"fmt"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/models"
	"log"
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

func TestUserValidation(t *testing.T) {
	user1 := &models.User{
		Username: "test",
		Password: "Phani",
	}

	user2 := &models.User{
		Username: "te",
		Password: "PhaniPhani",
	}

	user3 := &models.User{
		Username: "3test",
		Password: "3PhaniPhani",
	}
	user4 := &models.User{
		Username: "test(",
		Password: "PhaniPhani",
	}
	user5 := &models.User{
		Username: "test",
		Password: "PhaniPhani",
	}

	// first 4 test cases should fail and remaining to be passed
	_, err1 := user1.SignInValidation()
	_, err2 := user2.SignInValidation()
	_, err3 := user3.SignInValidation()
	_, err4 := user4.SignInValidation()
	_, err5 := user5.SignInValidation()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 == nil {
		log.Println("Validation was success full")
	} else {
		log.Println("err1", err1)
		log.Println("err2", err2)
		log.Println("err3", err3)
		log.Println("err4", err4)
		log.Println("err5", err5)
		t.Errorf("User validation failed")
	}
}

func TestSignUpValidation(t *testing.T) {
	// here we are checking only mail id and mobile number
	signIn := models.User{
		Username: "test",
		Password: "PhaniPhani",
	}
	user1 := &models.NewUser{
		User:   signIn,
		MailId: "phani<.123",
		Mobile: "123123123",
	}

	user2 := &models.NewUser{
		User:   signIn,
		MailId: "phani.123.com",
		Mobile: "123123123",
	}

	user3 := &models.NewUser{
		User:   signIn,
		MailId: "phani@123.com",
		Mobile: "123123123",
	}
	user4 := &models.NewUser{
		User:   signIn,
		MailId: "phani@MIL.com",
		Mobile: "123123123",
	}
	user5 := &models.NewUser{
		User:   signIn,
		MailId: "phani@mail.com",
		Mobile: "+11123123123",
	}

	// first 4 test cases should fail and remaining to be passed
	_, err1 := user1.SignUpValidation()
	_, err2 := user2.SignUpValidation()
	_, err3 := user3.SignUpValidation()
	_, err4 := user4.SignUpValidation()
	_, err5 := user5.SignUpValidation()
	if err1 != nil && err2 != nil && err3 != nil && err4 != nil && err5 == nil {
		log.Println("Validation was success full")
	} else {
		log.Println("err1", err1)
		log.Println("err2", err2)
		log.Println("err3", err3)
		log.Println("err4", err4)
		log.Println("err5", err5)
		t.Errorf("User validation failed")
	}
}

package models

type NewUser struct {
	username string
	password string
	mailId   string
	mobile   string
}

type User struct {
	username string
	password string
}

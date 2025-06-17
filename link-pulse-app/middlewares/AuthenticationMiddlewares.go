package middlewares

func AuthorizationCheck(authorizationHeader string) bool {
	return true
}

func CreateAuthorizationHeader(userId int32, username string) string {
	return "authorization header"
}

func MailIdCheck(mailId string) bool {
	return true
}

func MobileCheck(mobile string) bool {
	return true
}

func PasswordCheck(password string) bool {
	return true
}

func UsernameCheck(username string) bool {
	return true
}

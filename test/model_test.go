package auth_test

import "errors"

const (
	CORRECT_RESULT = "secret message"
)

var (
	ERR_INCORRECT_PASSWORD = errors.New("incorrect password")
	ERR_FORBIDDEN          = errors.New("forbidden")
	ERR_USER_NOT_FOUND     = errors.New("user not found")
)

type User struct {
	Username string
	Role     string
	Password string
}

type Login struct {
	Username string
	Password string
}

type TestModel struct {
	Name string
	*User
	Token string
	*Login
	Result error
}

var testModels = []*TestModel{
	&TestModel{
		Name:   "Success Case",
		User:   &User{Username: "lpteacher1", Password: "abc123", Role: "TEACHER"},
		Token:  "",
		Login:  &Login{Username: "lpteacher1", Password: "abc123"},
		Result: nil,
	},
	&TestModel{
		Name:   "Incorrect Login",
		User:   &User{Username: "lpstudent1", Password: "def456", Role: "STUDENT"},
		Token:  "",
		Login:  &Login{Username: "lpstudent1", Password: "def123"},
		Result: ERR_INCORRECT_PASSWORD,
	},
	&TestModel{
		Name:   "Forbidden Role",
		User:   &User{Username: "lpstudent2", Password: "ghi789", Role: "STUDENT"},
		Token:  "",
		Login:  &Login{Username: "lpstudent2", Password: "ghi789"},
		Result: ERR_FORBIDDEN,
	},
	&TestModel{
		Name:   "User Not Found",
		User:   &User{Username: "lpstudent2", Password: "ghi789", Role: "STUDENT"},
		Token:  "",
		Login:  &Login{Username: "lpstudent3", Password: "ghi789"},
		Result: ERR_USER_NOT_FOUND,
	},
}

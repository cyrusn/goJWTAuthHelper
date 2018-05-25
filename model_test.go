package auth_test

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
	*Login
}

var testModels = []*TestModel{
	&TestModel{
		Name: "Success Case",
		User: &User{
			Username: "lpteacher1",
			Password: "abc123",
			Role:     "TEACHER",
		},
		Login: &Login{
			Username: "lpteacher1",
			Password: "abc123",
		},
	},
	&TestModel{
		Name: "Incorrect Login",
		User: &User{
			Username: "lpstudent1",
			Password: "def456",
			Role:     "STUDENT",
		},
		Login: &Login{
			Username: "lpstudent1",
			Password: "def123",
		},
	},
	&TestModel{
		Name: "Forbidden Role",
		User: &User{
			Username: "lpstudent2",
			Password: "ghi789",
			Role:     "STUDENT",
		},
		Login: &Login{
			Username: "lpstudent2",
			Password: "ghi789",
		},
	},
}

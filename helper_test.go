package auth_test

import "errors"

func (m *TestModel) authenticate(username, password string) error {
	for _, m := range testModels {
		if m.User.Username == username && m.User.Password == password {
			return nil
		}
	}
	return errors.New("Unauthorised")
}

func in(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

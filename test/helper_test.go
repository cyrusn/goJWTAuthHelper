package auth_test

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func authAndCreateToken(username, password string) (*TestModel, error) {
	m, err := findAndAuth(username, password)
	if err != nil {
		return nil, err
	}

	token, err := name.CreateToken(myClaims{
		Username: m.User.Username,
		Role:     m.User.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	})

	if err != nil {
		return nil, err
	}

	m.Token = token
	return m, nil

}

func findAndAuth(username, password string) (*TestModel, error) {
	m, err := findModel(username)
	if err != nil {
		return nil, err
	}

	if err := m.authenticate(password); err != nil {
		return nil, err
	}

	return m, nil
}

func findModel(username string) (*TestModel, error) {
	for _, m := range testModels {
		if username == m.User.Username {
			return m, nil
		}
	}
	return nil, ERR_USER_NOT_FOUND
}

// Authenticate authorised user login
func (m *TestModel) authenticate(password string) error {
	if m.User.Password == password {
		return nil
	}
	return ERR_INCORRECT_PASSWORD
}

func in(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

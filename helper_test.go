package auth_test

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func authAndCreateToken(username, password string) (*TestModel, error) {
	m, err := findAndAuth(username, password)
	if err != nil {
		return nil, err
	}

	role := m.User.Role

	token, err := secret.CreateToken(myClaims{
		Username: username,
		Role:     role,
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

	if err := m.comparePassword(password); err != nil {
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

// comparePassword authorised user login
func (m *TestModel) comparePassword(password string) error {
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

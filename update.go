package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Updater is jwt.Claims but with update method to update the claims value
type Updater interface {
	Update(*jwt.Token)
	Valid() error
}

// UpdateToken update valid token
func (s *Secret) UpdateToken(tokenString string, r Updater) (string, error) {
	if tokenString == "" {
		return "", ErrTokenNotFound
	}
	token, err := jwt.ParseWithClaims(tokenString, r, s.keyFunc)
	if err != nil {
		return "", nil
	}

	r.Update(token)
	return token.SignedString(s.privateKey)
}

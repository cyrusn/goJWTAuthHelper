package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken return jwt token if the auth success, where claim will be stored
// in jwt payload
func (s *Secret) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.privateKey)
}

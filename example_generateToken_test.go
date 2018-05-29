package auth_test

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// ExampleSecret_GenerateToken shows how to create JWT token
func ExampleSecret_GenerateToken(username string, role string) (string, error) {
	return secret.GenerateToken(myClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAfter(30),
		},
	})
}

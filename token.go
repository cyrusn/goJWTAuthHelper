package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Authenticator is an interface that auth func
type Authenticator interface {
	Authenticate(loginName, password string) error
}

// CreateToken return jwt token if the auth success, where claim will be stored
// in jwt payload
func CreateToken(claim jwt.Claims, auth Authenticator, loginName, password string) (string, error) {
	if err := auth.Authenticate(loginName, password); err != nil {
		return "", err
	}

	return generateJWTToken(claim, privateKey)
}

func generateJWTToken(claim jwt.Claims, privateKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(privateKey)
}

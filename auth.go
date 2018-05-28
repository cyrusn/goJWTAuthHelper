// Package auth is a auth server package which will handle the login request
// user will receive a JWT Token once sucessfull login. JWT token will
// contain basic information of the user which depend on application call.
package auth

import "errors"

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrAccessDenied    = errors.New("accesss denied")
	ErrContextNotFound = errors.New("context not found")
	ErrRoleNotFound    = errors.New("role not found")
	ErrTokenNotFound   = errors.New("token not found")
)

// Secret store the information for Secret for auth package
type Secret struct {
	ContextKeyName string // the http.Context which store information jwt.Claims
	JWTKeyName     string
	RoleKeyName    string
	privateKey     []byte
}

// New return new key names
func New(ContextKeyName, JWTKeyName, RoleKeyName string, privateKey []byte) Secret {
	return Secret{ContextKeyName, JWTKeyName, RoleKeyName, privateKey}
}

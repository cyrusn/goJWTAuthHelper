// Package auth is a package which handles the authentication.
// User can receive a JWT string if user login sucessfully.
// JWT contains the information in payload session, user can define which
// information to be stored there.
//
// If User want to store information in payload session of jwt,
// You may customize a jwt.Claims. Please see example of create jwt.Claims
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
	// ContextKeyName is the key name of your custom jwt.Claims
	// which is stored in the http.Context.
	ContextKeyName string

	// JWTKeyName is the key name where the jwt be stored in the
	// HTTP headers of a request
	JWTKeyName string

	// RoleKeyName is a key name of role in your custom jwt.Claims
	RoleKeyName string

	// privateKey is private key of jwt
	privateKey []byte
}

// New return new key names
func New(ContextKeyName, JWTKeyName, RoleKeyName string, privateKey []byte) Secret {
	return Secret{ContextKeyName, JWTKeyName, RoleKeyName, privateKey}
}

// Package auth is a auth server package which will handle the login request
// user will receive a JWT Token once sucessfull login. JWT token will
// contain basic information of the user which depend on application call.
package auth

// Secret store the information for Secret for auth package
type Secret struct {
	contextKeyName string // the http.Context which store information jwt.Claims
	jwtKeyName     string
	roleKeyName    string
	privateKey     []byte
}

// New return new key names
func New(contextKeyName, jwtKeyName, roleKeyName string, privateKey []byte) Secret {
	return Secret{contextKeyName, jwtKeyName, roleKeyName, privateKey}
}

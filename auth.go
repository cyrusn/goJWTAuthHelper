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

// SetPrivateKey set the privateKey for authentication
// the default value of privateKey is secret
func (s *Secret) SetPrivateKey(key string) {
	s.privateKey = []byte(key)
}

// SetContextKeyName set the name of ContextKeyNameClaim,
// the default value of contextKeyName is "context-claims"
func (s *Secret) SetContextKeyName(name string) {
	s.contextKeyName = name
}

// SetJWTKeyName set the name of Role in jwt payload,
// the default value of roleKeyName is "jwt-token"
func (s *Secret) SetJWTKeyName(name string) {
	s.jwtKeyName = name
}

// SetRoleKeyName declares rolekeyName in jwt.Claims. auth package will
// get the value with the key named by rolekeyName to valdate the scope
// of authentication, the default value of role is "Role"
func (s *Secret) SetRoleKeyName(name string) {
	s.roleKeyName = name
}

// GetContextKeyName get the name of Role in jwt payload,
func (s *Secret) GetContextKeyName() string {
	return s.contextKeyName
}

// GetRoleKeyName get the name of Role in jwt payload,
func (s *Secret) GetRoleKeyName() string {
	return s.roleKeyName
}

// GetJWTKeyName get the name of Role in jwt payload,
func (s *Secret) GetJWTKeyName() string {
	return s.jwtKeyName
}

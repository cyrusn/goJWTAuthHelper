package auth

// Name store the names for auth package
type Name struct {
	contextKey string // the key in http.Context which store information jwt.Claims
	jwtKey     string
	roleKey    string
	privateKey []byte
}

// New return new key names
func New(contextKey, jwtKey, roleKey string, privateKey []byte) Name {
	return Name{contextKey, jwtKey, roleKey, privateKey}
}

// SetPrivateKey set the privateKey for authentication
// the default value of privateKey is secret
func (n *Name) SetPrivateKey(key string) {
	n.privateKey = []byte(key)
}

// SetContextKey set the name of ContextKeyClaim,
// the default value of contextKey is "context-claims"
func (n *Name) SetContextKey(name string) {
	n.contextKey = name
}

// SetJWTKey set the name of Role in jwt payload,
// the default value of roleKey is "jwt-token"
func (n *Name) SetJWTKey(name string) {
	n.jwtKey = name
}

// SetRoleKey declares rolekey in jwt.Claims. auth package will
// get the value with the key named by rolekey to valdate the scope
// of authentication, the default value of role is "Role"
func (n *Name) SetRoleKey(name string) {
	n.roleKey = name
}

// GetContextKey get the name of Role in jwt payload,
func (n *Name) GetContextKey() string {
	return n.contextKey
}

// GetRoleKey get the name of Role in jwt payload,
func (n *Name) GetRoleKey() string {
	return n.roleKey
}

// GetJWTKey get the name of Role in jwt payload,
func (n *Name) GetJWTKey() string {
	return n.jwtKey
}

package auth_test

import (
	"net/http"

	auth "github.com/cyrusn/goJWTAuthHelper"
	jwt "github.com/dgrijalva/jwt-go"
)

// Example show how to use Scope and Validate func
func Example() {
	secret = auth.New("myClaim", "kid", "myRole", []byte("secret"))
	for _, ro := range testRoutes {
		handler := http.HandlerFunc(ro.handler)

		if len(ro.scopes) != 0 {
			handler = secret.Scope(ro.scopes, handler).(http.HandlerFunc)
		}

		if ro.auth {
			handler = secret.Validate(handler).(http.HandlerFunc)
		}
		r.Handle(ro.path, handler)
	}
}

// ExampleSecret_CreateToken shows how to create JWT token
func ExampleSecret_CreateToken(username string, role string) (string, error) {
	return secret.CreateToken(myClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	})
}

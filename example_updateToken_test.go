package auth_test

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var lifeTime int64 = 30

func expiresAfter(lifeTime int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(lifeTime)).Unix()
}

type myClaims struct {
	Username string
	Role     string `json:"myRole"`
	jwt.StandardClaims
}

func (claims *myClaims) Update(token *jwt.Token) {
	claims.ExpiresAt = expiresAfter(lifeTime)
	token.Claims = claims
}

// ExampleSecret_UpdateToken is an example how to use UpdateToken func to
// update token
func ExampleSecret_UpdateToken() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token value in header with key "jwt-token"
		token := r.Header.Get("jwt-token")
		claims := myClaims{}
		newToken, err := secret.UpdateToken(token, &claims)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(newToken))
	}
}

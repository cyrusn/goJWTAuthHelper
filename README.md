# goJWTAuthHelper

Helper func for JWT Authentication of Golang

## Examples

```go
package auth_test

import (
	"errors"
	"net/http"
	"time"

	auth "github.com/cyrusn/goJWTAuthHelper"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	PUBLIC_CONTENT = "This is a public message."
	SECRET_CONTENT = "This is a secret message."
)

var (
	r             = mux.NewRouter()
	secret        = auth.New("myClaim", "kid", "myRole", []byte("secret"))
	expiresAfter  = 30
	ERR_FORBIDDEN = errors.New("forbidden")
)

type myClaims struct {
	Username string
	Role     string `json:"myRole"`
	jwt.StandardClaims
}

func expiresAt() int64 {
	return time.Now().Add(time.Minute * time.Duration(expiresAfter)).Unix()
}

func (claims *myClaims) Update(token *jwt.Token) {
	claims.ExpiresAt = expiresAt()
	token.Claims = claims
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	errCode := http.StatusForbidden

	username := mux.Vars(r)["username"]
	role := mux.Vars(r)["role"]

	authorized := true
	if !authorized {
		http.Error(w, ERR_FORBIDDEN.Error(), errCode)
		return
	}

	token, err := secret.GenerateToken(myClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt(),
		},
	})

	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}

	w.Write([]byte(token))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	// Get jwt string in header with key "kid"
	token := r.Header.Get("kid")
	newToken, err := secret.UpdateToken(token, new(myClaims))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Write([]byte(newToken))
}

type route struct {
	path    string
	auth    bool
	scopes  []string
	methods []string
	handler func(http.ResponseWriter, *http.Request)
}

var routes = []route{
	route{
		path:    "/login/{username}/{role}",
		auth:    false,
		scopes:  []string{},
		methods: []string{"POST"},
		handler: loginHandler,
	},
	route{
		path:    "/refresh",
		auth:    true,
		scopes:  []string{},
		methods: []string{"GET"},
		handler: refreshHandler,
	},
	route{
		path:    "/secret",
		auth:    true,
		scopes:  []string{"teacher"},
		methods: []string{"GET"},
		handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(SECRET_CONTENT))
		},
	},
	route{
		path:    "/public",
		auth:    false,
		scopes:  []string{},
		methods: []string{"GET"},
		handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(PUBLIC_CONTENT))
		},
	},
}

func Example() {
	for _, ro := range routes {
		handler := http.HandlerFunc(ro.handler)

		// pass Access to handler first
		if len(ro.scopes) != 0 {
			handler = secret.Access(ro.scopes, handler).(http.HandlerFunc)
		}

		// then pass Authenticate at last
		if ro.auth {
			handler = secret.Authenticate(handler).(http.HandlerFunc)
		}

		r.
			Methods(ro.methods...).
			Path(ro.path).
			HandlerFunc(handler)
	}
}
```
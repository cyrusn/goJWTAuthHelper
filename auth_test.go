package auth_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/cyrusn/goJWTAuthHelper"
	"github.com/cyrusn/goTestHelper"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var (
	name            = auth.New("myClaim", "kid", "myRole", []byte("secret"))
	r               = mux.NewRouter()
	expireToken     = time.Now().Add(time.Minute * 30).Unix()
	authorizedRoles = []string{"TEACHER"}
	token           = ""
)

func init() {
	for _, ro := range testRoutes {
		handler := http.HandlerFunc(ro.handler)

		if len(ro.scopes) != 0 {
			handler = name.Scope(ro.scopes, handler).(http.HandlerFunc)
		}

		if ro.auth {
			handler = name.Validate(handler).(http.HandlerFunc)
		}
		r.Handle(ro.path, handler)
	}
}

func TestMain(t *testing.T) {
	t.Run("Login:", loginAndGotoAuth)
	t.Run("/basic/", basicRouteTest)
}

type myClaims struct {
	Username string
	Role     string `json:"myRole"`
	jwt.StandardClaims
}

var loginAndGotoAuth = func(t *testing.T) {
	for _, m := range testModels {
		t.Run(m.Name, func(t *testing.T) {
			t.Run("/login/", loginTest(m))
			t.Run("/auth/", authRouteTest(m))
		})
	}
}

var basicRouteTest = func(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/basic/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.OK(t, err)
	assert.Equal(string(body), "secret message", t)
}

var loginTest = func(m *TestModel) func(*testing.T) {
	return func(t *testing.T) {
		loginWriter := httptest.NewRecorder()
		formString := fmt.Sprintf(`{"Username":"%s", "Password":"%s"}`, m.User.Username, m.User.Password)
		postForm := strings.NewReader(formString)
		loginReq := httptest.NewRequest("POST", "/login/", postForm)
		loginReq.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(loginWriter, loginReq)

		loginResp := loginWriter.Result()
		tokenBytes, err := ioutil.ReadAll(loginResp.Body)
		assert.OK(t, err)
		token = string(tokenBytes)
	}
}

var authRouteTest = func(m *TestModel) func(*testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/auth/", nil)
		req.Header.Set(name.GetJWTKey(), token)
		r.ServeHTTP(w, req)

		resp := w.Result()
		body, err := ioutil.ReadAll(resp.Body)
		assert.OK(t, err)

		username := m.User.Username
		password := m.User.Password
		correctScope := in(m.User.Role, authorizedRoles)

		if m.authenticate(username, password) == nil && correctScope {
			assert.Equal(string(body), "secret message", t)
		} else {
			assert.Panic(m.Name, t, func() {
				panic(string(body))
			})
		}
	}

}

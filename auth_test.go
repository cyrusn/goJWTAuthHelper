package auth_test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/cyrusn/goJWTAuthHelper"
	"github.com/cyrusn/goTestHelper"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var (
	secret          auth.Secret
	r               = mux.NewRouter()
	expireToken     = time.Now().Add(time.Minute * 30).Unix()
	authorizedRoles = []string{"TEACHER"}
)

func init() {
	Example()
}

func TestMain(t *testing.T) {
	t.Run("auth test", loginAndGotoAuthTest)
	t.Run("basic test", basicTest)
}

type myClaims struct {
	Username string
	Role     string `json:"myRole"`
	jwt.StandardClaims
}

var loginAndGotoAuthTest = func(t *testing.T) {
	for _, m := range testModels {
		m.Token = ""
		t.Run(m.Name, func(t *testing.T) {
			t.Run("login", loginTest(m))
			t.Run("auth", authRouteTest(m))
		})
	}
}

var basicTest = func(t *testing.T) {
	for _, m := range testModels {
		m.Token = ""
		t.Run(m.Name, func(t *testing.T) {
			t.Run("login", loginTest(m))
			t.Run("basic", func(t *testing.T) {
				w := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/basic/", nil)
				r.ServeHTTP(w, req)

				resp := w.Result()
				body, err := ioutil.ReadAll(resp.Body)
				assert.OK(t, err)
				assert.Equal(string(body), CORRECT_RESULT, t)
			})
		})
	}
}

var loginTest = func(m *TestModel) func(*testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		formJSON := fmt.Sprintf(`{"Username":"%s", "Password":"%s"}`, m.Login.Username, m.Login.Password)
		postForm := strings.NewReader(formJSON)
		req := httptest.NewRequest("POST", "/login/", postForm)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		resp := w.Result()
		body, err := ioutil.ReadAll(resp.Body)
		assert.OK(t, err)

		if resp.StatusCode < 400 {
			m.Token = string(body)
			return
		}

		if m.Result == nil {
			assert.Equal(string(body), CORRECT_RESULT, t)
		}

		errMsg, err := helper.UnmarshalErrMessage(body)
		assert.OK(t, err)

		errs := []error{
			ERR_FORBIDDEN,
			ERR_INCORRECT_PASSWORD,
			ERR_USER_NOT_FOUND,
		}

		for _, e := range errs {
			if m.Result == e {
				assert.Equal(errMsg.Message, e.Error(), t)
			}
		}

	}
}

var authRouteTest = func(m *TestModel) func(*testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/auth/", nil)
		req.Header.Set("kid", m.Token)
		r.ServeHTTP(w, req)

		resp := w.Result()
		body, err := ioutil.ReadAll(resp.Body)
		if m.Result == nil {
			assert.OK(t, err)
			assert.Equal(string(body), CORRECT_RESULT, t)
		}

		// check invalid access
		assert.Panic(m.Name, t, func() {
			panic(string(body))
		})
	}

}

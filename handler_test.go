package auth_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("secret message"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	type LoginForm struct {
		Username string
		Password string
	}
	errCode := http.StatusUnauthorized
	loginForm := new(LoginForm)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	if err := json.Unmarshal(body, loginForm); err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	username := loginForm.Username
	password := loginForm.Password

	m, err := authAndGenerateToken(username, password)
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	w.Write([]byte(m.Token))
}

package auth_test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyrusn/goTestHelper"
)

var (
	authorizedRoles = []string{"TEACHER"}
	jwtToken        = ""
	startTime       = time.Now()
)

func TestMain(t *testing.T) {
	Example()

	t.Run("public test", publicTest)
	t.Run("fail access test", failAccessTest(403))
	t.Run("login test", loginTest("teacher1", "teacher"))
	t.Run("success access test", successAccessTest)
	t.Run("login test", loginTest("student1", "student"))
	t.Run("fail access test", failAccessTest(401))
	t.Run("Refresh token", refreshTest)
}

var publicTest = func(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/public", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.OK(t, err)
	assert.Equal(string(body), PUBLIC_CONTENT, t)
}

func failAccessTest(statusCode int) func(*testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/secret", nil)
		req.Header.Set("kid", jwtToken)
		r.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(resp.StatusCode, statusCode, t)
	}
}

var loginTest = func(username, role string) func(*testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		path := fmt.Sprintf("/login/%s/%s", username, role)
		req := httptest.NewRequest("POST", path, nil)
		r.ServeHTTP(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode < 300 {
			jwtToken = string(body)
			assert.OK(t, nil)
		}
	}
}

func successAccessTest(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/secret", nil)
	req.Header.Set("kid", jwtToken)
	r.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(resp.StatusCode, 200, t)

	body, _ := ioutil.ReadAll(resp.Body)
	got := string(body)
	assert.Equal(got, SECRET_CONTENT, t)
}

var refreshTest = func(t *testing.T) {
	time.Sleep(5000 * time.Millisecond)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/refresh", nil)
	req.Header.Set("kid", jwtToken)
	r.ServeHTTP(w, req)
	resp := w.Result()
	assert.Equal(resp.StatusCode, 200, t)

	// Uncomments the following lines to get the refreshed JWT,
	// you may use the tool in https://jwt.io/#debugger to see decode the token

	// fmt.Println(startTime.
	// 	Add(time.Minute * time.Duration(expiresAfter)).
	// 	Add(time.Second * 5))
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(jwtToken)
	// fmt.Println(string(body))
}

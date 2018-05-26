package auth_test

import "net/http"

type route struct {
	path    string
	auth    bool
	scopes  []string
	method  string
	handler func(http.ResponseWriter, *http.Request)
}

var testRoutes = []route{
	route{
		path:    "/login/",
		auth:    false,
		scopes:  []string{},
		method:  "GET",
		handler: loginHandler,
	},
	route{
		path:    "/auth/",
		auth:    true,
		scopes:  authorizedRoles,
		method:  "GET",
		handler: simpleHandler,
	},
	route{
		path:    "/basic/",
		auth:    false,
		scopes:  []string{},
		method:  "GET",
		handler: simpleHandler,
	},
}

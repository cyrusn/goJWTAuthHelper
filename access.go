package auth

import (
	"context"
	"net/http"

	"github.com/cyrusn/goHTTPHelper"

	jwt "github.com/dgrijalva/jwt-go"
)

// Access is a middleware that parse jwt in header with value of roleKeyName as key
// (default value of roleKeyName is "Role", user can use SetRoleKeyName to set
// the vale of the roleKeyName).
// If value with roleKeyName in jwt payload is not in "scopes []string", handler will
// print error message instead
func (s *Secret) Access(scopes []string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusUnauthorized
		role, err := s.parseRoleInContext(r.Context())
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		if ok := in(role, scopes); !ok {
			helper.PrintError(w, ErrAccessDenied, errCode)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (s *Secret) parseRoleInContext(ctx context.Context) (string, error) {
	contextKeyName := s.contextKeyName
	roleKeyName := s.roleKeyName

	claim := ctx.Value(contextKeyName)

	if claim == nil {
		return "", ErrContextNotFound
	}
	m := claim.(jwt.MapClaims)
	result, ok := m[roleKeyName].(string)
	if !ok {
		return "", ErrRoleNotFound
	}

	return result, nil
}

func in(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

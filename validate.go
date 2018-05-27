package auth

import (
	"context"
	"net/http"

	"github.com/cyrusn/goHTTPHelper"
	jwt "github.com/dgrijalva/jwt-go"
)

// Authenticate is a middleware which will check if jwt in request header is valid
func (s *Secret) Authenticate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(s.jwtKeyName)
		if jwtToken == "" {
			errCode := http.StatusForbidden
			helper.PrintError(w, ErrTokenNotFound, errCode)
			return
		}

		token, err := jwt.Parse(jwtToken, s.keyFunc)
		errCode := http.StatusUnauthorized
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), s.contextKeyName, claims)
			req := r.WithContext(ctx)
			handler.ServeHTTP(w, req)
			return
		}
		helper.PrintError(w, ErrInvalidToken, errCode)
	})
}

func (s *Secret) keyFunc(token *jwt.Token) (interface{}, error) {
	return s.privateKey, nil
}

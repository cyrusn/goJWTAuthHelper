package auth_test

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func expire(min int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(min)).Unix()
}

type myClaims struct {
	Username string
	Role     string `json:"myRole"`
	jwt.StandardClaims
}

func (claims *myClaims) Update(token *jwt.Token) {
	claims.ExpiresAt = expire(10)
	token.Claims = claims
}

func ExampleSecret_UpdateToken() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImxwdGVhY2hlcjEiLCJteVJvbGUiOiJURUFDSEVSIiwiZXhwIjoxNTI3NDM3MjQ2fQ.qtTwRKJ5YiPkRhVxFp2-yQBwydMibOpnr685u7X7jC0"
	claims := myClaims{}
	newToken, err := secret.UpdateToken(token, &claims)
	if err != nil {
		panic(err)
	}
	fmt.Println(newToken)
}

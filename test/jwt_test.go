package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/jwt"
	"testing"
)

func TestJwtCreate(t *testing.T) {
	t.Helper()

	token := jwt.Jwt{
		Payload: &jwt.Payload{
			Subject:   "",
			Issuer:    "",
			ExpiresAt: 100,
			//Scope:     []string{"user"},
		},
		Secret: []byte("test123"),
	}
	tokenStr, err := token.Create()
	if err != nil {
		t.Errorf("create token failed，err：%s", err.Error())
	}

	fmt.Println(tokenStr)
	claims, err := token.Parse(tokenStr)
	if err != nil {
		t.Errorf("parse token failed, err: %s", err.Error())
	}
	fmt.Println(claims)
}

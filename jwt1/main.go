package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	// example code for custom claims
	// https://godoc.org/github.com/dgrijalva/jwt-go#example-NewWithClaims--CustomClaimsType
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v\n", ss, err)
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	tokenString := generate()
	fmt.Println("Generated token:", tokenString)

	{
		c1 := parse(tokenString)
		fmt.Printf("Parsed IssuedAt: %#v\n", c1.IssuedAt)
	}

	fmt.Println(time.Local) // Call time.Local

	{
		c2 := parse(tokenString)
		fmt.Printf("Parsed IssuedAt: %#v\n", c2.IssuedAt)
	}
}

func generate() string {
	t := time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC)
	claims := jwt.RegisteredClaims{
		IssuedAt: jwt.NewNumericDate(t),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func parse(tokenString string) *jwt.RegisteredClaims {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &jwt.RegisteredClaims{})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		log.Fatal("failed to parse claims")
	}

	return claims
}

package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"fmt"

	"strconv"
)

type authJsonResponse struct {
	Stat        string `json:"status"`
	TokenString string `json:"tokenString"`
}

type MyCustomClaims struct {
	Remark string `json:"remark"`
	Level  int    `json:"level"`
	jwt.RegisteredClaims
}

func jwtTest() {
	mySigningKey := getSigningKey()
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
	}
	fmt.Println(claims, mySigningKey)
}

func getSigningKey() []byte {
	mySigningKey := []byte(ENV_INPUTS.Secrets.JwtKey)
	return mySigningKey
}

func generateTokenString(user userDetails) string {

	mySigningKey := getSigningKey()

	// Create the claims
	claims := MyCustomClaims{
		"Auth Registered",
		user.AccessLevel,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "QdasServicePortal",
			Subject:   "authToken",
			ID:        strconv.FormatInt(int64(user.ID), 10),
			Audience:  []string{"authenticated_user"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
	}
	return ss
}

func validateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return getSigningKey(), nil
		},
	)
	if err != nil {
		return false
	}
	if token.Valid {
		return true
	}
	return false
}

func getClaimsFromValidToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return getSigningKey(), nil
		},
	)
	var abc *MyCustomClaims
	if err != nil {
		return abc, err
	}
	claims := token.Claims.(*MyCustomClaims)
	return claims, nil

}

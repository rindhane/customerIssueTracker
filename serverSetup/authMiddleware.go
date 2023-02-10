package main

import (
	"context"

	"github.com/gin-gonic/gin"
)

var contextAuthKey string = "tokenString"
var authTokenKey string = "authToken"
var validAuthKeyInContext = "ValidAuth"

func updateContextWithAuth(c *gin.Context, tokenString string) {
	cxt := c.Request.Context()
	newCxt := context.WithValue(cxt, validAuthKeyInContext, true)
	newCxt = context.WithValue(newCxt, contextAuthKey, tokenString)
	c.Request = c.Request.Clone(newCxt)
	return
}

func getClaimsFromAuthToken(c *gin.Context) (*MyCustomClaims, error) {
	//it is assumed that authToken has already been validated
	value, _ := c.Request.Context().Value(contextAuthKey).(string)
	claims, err := getClaimsFromValidToken(value)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func AuthMiddleWare(c *gin.Context) {
	//check whether valid token provided in header
	HeaderTokenString := c.GetHeader(authTokenKey)
	if HeaderTokenString != "" {
		if validateToken(HeaderTokenString) {
			updateContextWithAuth(c, HeaderTokenString)
			c.Next()
			return
		} else {
			c.Next()
			return
		}
	}
	//check whether valid cookie provided in request
	tokenString, _ := c.Cookie(authTokenKey)
	if validateToken(tokenString) {
		updateContextWithAuth(c, tokenString)
		c.Next()
		return
	}
	c.Next()
	return
}

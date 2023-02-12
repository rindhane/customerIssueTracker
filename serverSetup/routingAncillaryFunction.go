package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func authTest(c *gin.Context) bool {
	ctx := c.Request.Context()
	value := ctx.Value("ValidAuth")
	if value == true {
		return true
	}
	return false
}

func authFailAction(c *gin.Context) {
	//authentication set
	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
	return
	//redirect with params if needed, ref:https://stackoverflow.com/questions/61970551/golang-gin-redirect-and-render-a-template-with-new-variables
}

func HomePage(c *gin.Context) {
	location := url.URL{Path: "/UserDashboard"}
	c.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
}

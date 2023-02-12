package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type issueObj struct {
	Id         string `json:"issueID"`
	ShortDesc  string `json:"shortDesc"`
	LastAction string `json:"lastAction"`
	Status     string `json:"status"`
}

type resultStatus struct {
	Status string `json:"status"`
	Remark string `json:"remark"`
}

type loginCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var sample = []issueObj{
	{Id: "1", ShortDesc: "company1-Issue:1", LastAction: "Ishwar", Status: "Assigned"},
	{Id: "2", ShortDesc: "company1-Issue:2", LastAction: "Virendra", Status: "Initiated"},
}

func loginPage(c *gin.Context) {
	var pageName string = "login.html"
	c.HTML(
		http.StatusOK,
		pageName,
		gin.H{
			"title": "Login",
		},
	)
}

// function to provide auth token if the user credentials are satisfied
func (ct *Controller) checkAuth(c *gin.Context) {
	credFromClient, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: "not authenticated"})
		return
	}
	//unwrap credentials provided from json string
	userCred := loginCredential{}
	json.Unmarshal(credFromClient, &userCred)
	//validate UserCredentials
	valid, user := validateUserCredentials(c, ct, userCred.Email, userCred.Password)
	if !valid {
		c.JSON(http.StatusUnauthorized, resultStatus{
			Status: "false",
			Remark: "not authenticated",
		})
		return
	}
	auth := authJsonResponse{
		Stat:        "true",
		TokenString: generateTokenString(userDetails{ID: user.ID}),
	}
	c.JSON(http.StatusOK, auth)
	return
}

func getUserIssues(c *gin.Context) {
	if !authTest(c) {
		authFailAction(c)
		return
	}
	c.JSON(http.StatusOK, sample)
	value, exists := c.Get("data123")
	if exists {
		fmt.Println("something:", value)
	}
}

func getUserDashboard(c *gin.Context) {
	if !authTest(c) {
		authFailAction(c)
		return
	}

	var name string = "UserDashboard.html"
	c.HTML(
		http.StatusOK,
		name,
		gin.H{
			"title":   "HomePage",
			"year":    "2023",
			"company": "Hexagon Manufacturing Intelligence",
		},
	)
}

func (ct *Controller) newIssueRaise(c *gin.Context) {
	if !authTest(c) {
		authFailAction(c)
		return
	}
	issueFromClient, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: "something went wrong"})
		return
	}
	issue := issueDetails{}
	json.Unmarshal(issueFromClient, &issue)
	claims, err := getClaimsFromAuthToken(c)
	if err == nil {
		userId, err := strconv.Atoi(claims.ID)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: "something went wrong"})
			return
		}
		issue.User = userId
		issue.Status = 1
		if enterNewIssue(c.Request.Context(), ct, issue) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"remark": "issue recorded by the server",
			})
			return
		}
	}
	c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: "something went wrong"})
	return
}

func fetchOTPRequest(c *gin.Context) {
	credFromClient, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: "no input"})
		return
	}
	userCred := loginCredential{}
	json.Unmarshal(credFromClient, &userCred)
	err = initiateOTP(userCred.Email)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"remark": "OTP dispatched over email. OTP is valid for 5 mins",
		})
		return
	}
	c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: err.Error()})
}

func (ct *Controller) otpAuthValidationSignUpReset(c *gin.Context) {
	credFromClient, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: "no input"})
		return
	}
	userCred := loginCredential{}
	json.Unmarshal(credFromClient, &userCred)
	err = validateOTPNewAccount(ct, c.Copy(), userCred.Email, userCred.Password)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, resultStatus{Status: "error", Remark: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultStatus{
		Status: "ok",
		Remark: "Account has been updated. New Password will sent on the email for login",
	})
	fmt.Println("validation of OTP", userCred)
	return
}

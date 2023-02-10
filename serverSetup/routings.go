package main

import (
	"github.com/gin-gonic/gin"
)

func setRoutes(r *gin.Engine, ct Controller) {
	r.POST("/UserIssues", getUserIssues)
	r.GET("/UserDashboard", getUserDashboard)
	r.GET("/login", loginPage)
	r.POST("/checkAuth", checkAuth)
	r.POST("/newIssueRaise", ct.newIssueRaise)
	r.POST("/generateOTP", fetchOTPRequest)
	r.POST("/OtpAuth", otpAuthValidation)
}

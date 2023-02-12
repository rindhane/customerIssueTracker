package main

import (
	"customerSite/serverSetup/azuredb"
	"fmt"
	Reader "helpers/tomlReader"

	"github.com/gin-gonic/gin"
)

type (
	tomlData struct {
		Secrets  keyConfig        `toml:"secrets"`
		DbConfig azuredb.DbConfig `toml:"dbConfig"`
	}

	keyConfig struct {
		OtpServiceEndpoint string `toml:"otpEndpoint"`
		JwtKey             string `toml:"jwtKey"`
	}
)

var ENV_INPUTS tomlData

func setEnvInputs(envInput *tomlData) {
	Reader.ReadTomlConfig("./secrets.toml", envInput)
}

func main() {
	setEnvInputs(&ENV_INPUTS)
	router := gin.Default()
	router.Use(AuthMiddleWare)
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	// set db setup
	//reference : https://stackoverflow.com/questions/35672842/go-and-gin-passing-around-struct-for-database-context
	pool, err := azuredb.InstantiateDBpool(ENV_INPUTS.DbConfig.GetConnString())
	defer pool.Close()
	if err != nil {
		fmt.Println("db correction error occured: ", err.Error())
	}
	ct := Controller{Database: pool} // preparing for passing db reference
	setRoutes(router, &ct)
	router.Run("localhost:9500")

}

//pending:
/*
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
*/

//complete from
//https://go.dev/doc/tutorial/web-service-gin
//https://www.scaler.com/event/golang-crash-course/
//https://www.scaler.com/events
//https://www.youtube.com/watch?v=yyUHQIec83I&t=8s
//https://betterprogramming.pub/how-to-render-html-pages-with-gin-for-golang-9cb9c8d7e7b6

//middleware for auth :
//https://sosedoff.com/2014/12/21/gin-middleware.html
//https://auth0.com/blog/authentication-in-golang/
//https://mattermost.com/blog/how-to-build-an-authentication-microservice-in-golang-from-scratch/

//other ref:
//https://faun.pub/dependency-injection-in-go-the-better-way-833e2cc65fab
//https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/

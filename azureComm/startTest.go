package main

import (
	"azureComm/azureOTP"
	"fmt"
	Reader "helpers/tomlReader"
)

type Addrs struct {
	Endpoint1 string `toml:"endpoint"`
}

type tomlConfig struct {
	AddrStruct Addrs `toml:"clientAddr"`
}

func main() {
	var EmailRequest azureOTP.EmailRequest
	EmailRequest.Email = "test@example.com"
	EmailRequest.Message = "123456"
	var test1 tomlConfig
	_, _ = Reader.ReadTomlConfig("./secret.toml", &test1)
	EmailRequest.ServiceUrl = test1.AddrStruct.Endpoint1
	result, err := azureOTP.SendOTPRequestAzure(&EmailRequest)
	if err != nil {
		fmt.Println("went wrong in sending otp", err)
	}
	fmt.Println(result)
}

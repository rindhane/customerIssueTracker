package main

import (
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
	var tomlStruct tomlConfig
	_, err := Reader.ReadTomlConfig("./secret.toml", &tomlStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("output 1 ", tomlStruct.AddrStruct)
}

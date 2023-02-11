package Reader

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func checkError(e error, print bool) error {
	if e != nil {
		if print {
			fmt.Println(e.Error())
		}
		return e
	}
	return nil
}

func getTomlString(filePath string) string {
	data, err := os.ReadFile(filePath)
	checkError(err, false)
	return string(data)
}

func ReadTomlConfig(tomlfilePath string, output any) (any, error) {
	tomlData := getTomlString(tomlfilePath)
	meta, err := toml.Decode(tomlData, output)
	if checkError(err, false) != nil {
		return output, err
	}
	return meta, nil
}

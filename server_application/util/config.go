package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Load(mode, srcPath string, cfg interface{}) error {
	configFileName := fmt.Sprintf("application.%s.json", mode)
	configFilePath := path.Join(srcPath, "config", configFileName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFilePath = path.Join(srcPath, "config", "application.json")
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &cfg)
}

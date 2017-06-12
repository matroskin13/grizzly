package gen

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type GrizzlyConfigCollection struct {
	Name string `json:"name"`
	Types map[string]string `json:"types"`
	Methods []string
}

type GrizzlyConfig struct {
	Collections []GrizzlyConfigCollection `json:"collections"`
}

func GetConfig() (config *GrizzlyConfig, err error) {
	currentPath, _ := os.Getwd();
	fullPwd := filepath.Join(currentPath, "grizzly.json")

	bytes, err := ioutil.ReadFile(fullPwd)

	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return config, err
	}

	return config, err
}

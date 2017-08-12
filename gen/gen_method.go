package gen

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetMethod(collectionDir string, method string) (result []byte, err error) {
	methodPath := filepath.Join(collectionDir, method+".go")

	if CheckExistFile(methodPath) {
		return result, fmt.Errorf("method %s is not exist", method)
	}

	result, err = ioutil.ReadFile(methodPath)

	if err != nil {
		return result, err
	}

	return result, err
}

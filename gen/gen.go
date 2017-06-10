package gen

import (
	"io/ioutil"
	"regexp"
	"os"
	"path/filepath"
	"strings"
	"errors"
	"fmt"
)

const GithubRepo = "github.com/matroskin13/grizzly"

func GetCollectionDir(isDev bool) (string, error) {
	goPaths := strings.Split(os.Getenv("GOPATH"), ":")

	if isDev {
		return "./collection/collection.go", nil
	}

	for _, path := range goPaths {
		grizzlyPath := filepath.Join(path, GithubRepo)

		fmt.Println("find grizzly in", grizzlyPath)

		if CheckExist(grizzlyPath) {
			return path, nil
		}
	}

	return "", errors.New("grizzly repo is not defined")
}

func GetCollectionCode(isDev bool, modelName string) (result string, err error) {
	collectionDir, err := GetCollectionDir(isDev)

	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(collectionDir)

	if err != nil {
		return result, err
	}

	rModel, _ := regexp.Compile("Model")
	code := rModel.ReplaceAll(bytes, []byte(modelName))

	rCollections, _ := regexp.Compile("Collection")
	code = rCollections.ReplaceAll(code, []byte(modelName + "Collection"))

	pCollections, _ := regexp.Compile("package collection")
	code = pCollections.ReplaceAll(code, []byte("package collections"))

	result = string(code)

	return result, err
}

func CheckExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil || os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CreateCollection(modelName string, code string) error {
	pwd, _ := os.Getwd()
	collectionPath := filepath.Join(pwd, "collections")
	filePath := filepath.Join(collectionPath, modelName + ".go");

	if !CheckExist(collectionPath) {
		os.Mkdir(collectionPath, os.ModePerm)
	}

	if !CheckExist(filePath) {
		err := ioutil.WriteFile(filePath, []byte(code), 0666)

		if err != nil {
			return err
		}
	}

	return nil
}
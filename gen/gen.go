package gen

import (
	"io/ioutil"
	"regexp"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

const GithubRepo = "github.com/matroskin13/grizzly"

func GetCollectionDir(isDev bool) (string, error) {
	goPaths := strings.Split(os.Getenv("GOPATH"), ":")

	if isDev {
		return "./collection/collection.go", nil
	}

	for _, path := range goPaths {
		grizzlyPath := filepath.Join(path, "src", GithubRepo)

		if !CheckExistDir(grizzlyPath) {
			return filepath.Join(grizzlyPath, "collection/collection.go"), nil
		}
	}

	return "", errors.New("grizzly repo is not defined")
}

func GetCollectionCode(isDev bool, modelName string, types map[string]string) (result string, err error) {
	collectionDir, err := GetCollectionDir(isDev)
	modelName = strings.Title(modelName)

	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(collectionDir)

	if err != nil {
		return result, err
	}

	var structString = " {\n"

	for key, value := range types {
		structString += "\t" + strings.Title(key) + " " + value + "\n"
 	}

	structString += "}";

	rStruct, _ := regexp.Compile("type Model struct \\{GrizzlyId int; GrizzlyName string}")
	code := rStruct.ReplaceAll(bytes, []byte("type Model struct" + structString))

	rSearchCallback, _ := regexp.Compile("SearchCallback")
	code = rSearchCallback.ReplaceAll(code, []byte("SearchCallback" + strings.Title(modelName)))

	rModel, _ := regexp.Compile("Model")
	code = rModel.ReplaceAll(code, []byte(modelName))

	rCollections, _ := regexp.Compile("Collection")
	code = rCollections.ReplaceAll(code, []byte(modelName + "Collection"))

	pCollections, _ := regexp.Compile("package collection")
	code = pCollections.ReplaceAll(code, []byte("package collections"))

	result = string(code)

	return result, err
}

func CheckExistDir(path string) bool {
	_, err := os.Stat(path)

	if err == nil || os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CheckExistFile(path string) bool {
	stat, _ := os.Stat(path)

	if stat == nil {
		return false
	} else {
		return true
	}
}

func CreateCollection(modelName string, code string, isUpdate bool) error {
	pwd, _ := os.Getwd()
	collectionPath := filepath.Join(pwd, "collections")
	filePath := filepath.Join(collectionPath, modelName + ".go");

	if !CheckExistDir(collectionPath) {
		os.Mkdir(collectionPath, os.ModePerm)
	}

	if !CheckExistFile(filePath) || isUpdate {
		err := ioutil.WriteFile(filePath, []byte(code), 0666)

		if err != nil {
			return err
		}
	} else {
		return errors.New("collection is alredy exist")
	}

	return nil
}
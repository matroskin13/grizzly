package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

const GithubRepo = "github.com/matroskin13/grizzly"

const (
	MethodFind = "find"
	MethodFilter = "filter"
	MethodMaps = "maps"
	MethodArray = "array"
	MethodGet = "get"
)

func GetDefaultMethods() []string {
	return []string {
		MethodFind,
		MethodFilter,
		MethodMaps,
		MethodArray,
		MethodGet,
	}
}

func GetCollectionDir(isDev bool) (string, error) {
	goPaths := strings.Split(os.Getenv("GOPATH"), ":")

	if isDev {
		return "./collection", nil
	}

	for _, path := range goPaths {
		grizzlyPath := filepath.Join(path, "src", GithubRepo)

		if !CheckExistDir(grizzlyPath) {
			return filepath.Join(grizzlyPath, "collection"), nil
		}
	}

	return "", errors.New("grizzly repo is not defined")
}

func GetCollectionCode() (result []byte, err error) {
	collectionDir, err := GetCollectionDir(false)

	if err != nil {
		return result, err
	}

	collectionPath := filepath.Join(collectionDir, "collection.go")

	result, err = ioutil.ReadFile(collectionPath)

	if err != nil {
		return result, err
	}

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

func GetMethodsCode(methods []string) (result []byte, err error) {
	collectionDir, err := GetCollectionDir(false)

	if err != nil {
		return result, err
	}

	for _, v := range methods {
		methodPath := filepath.Join(collectionDir, v + ".go")

		bytes, err := ioutil.ReadFile(methodPath)

		if err != nil {
			return result, err
		}

		result = append(result, bytes...)
	}

	return result, err
}

func GenCollectionCode(config GrizzlyConfigCollection) (result string, err error) {
	code, err := GetCollectionCode()

	if err != nil {
		return result, err
	}

	methodCode, err := GetMethodsCode(config.Methods)

	code = append(code, methodCode...)

	if err != nil {
		return result, err
	}

	code = ReplaceModel(code, config.Name, config.Types)
	code = ReplaceSearchCallback(code, config.Name)
	code = ReplaceCollection(code, config.Name)

	code = append([]byte("package collections"), code...)

	return string(code), err
}
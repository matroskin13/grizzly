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
	MethodUniq = "uniq"
	MethodSort = "sort"
	MethodEach = "each"
)

const (
	GrizzlyCollection = "Collection"
	GrizzlyModel = "Model"

	CommandReplaceName = "grizzly:replaceName"
	CommandGenerate    = "grizzly:generate"
)

type GrizzlyCommand struct {
	Command string
	Action string
}

func GetDefaultMethods() []string {
	return []string {
		MethodFind,
		MethodFilter,
		MethodMaps,
		MethodArray,
		MethodGet,
		MethodUniq,
		MethodSort,
		MethodEach,
	}
}

func IsPropertyMethod(method string) bool {
	for _, propertyMethod := range []string{MethodUniq, MethodSort} {
		if propertyMethod == method {
			return true
		}
	}

	return false
}

func GetImportsByMethods(methods []string) (imports []string) {
	for _, method := range methods {
		switch method {
		case MethodSort:
			imports = append(imports, "sort")
		}
	}

	return imports
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

	code := RemovePackage(result)

	return code, err
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

func CreateCollection(modelName string, code string, isUpdate bool, savePath string) error {
	var collectionPath string

	pwd, _ := os.Getwd()

	if savePath == "" {
		collectionPath = filepath.Join(pwd, "collections")
	} else {
		collectionPath = savePath
	}

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

func GetMethodsCode(methods []string, types []GrizzlyType) (result []byte, err error) {
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

		if IsPropertyMethod(v) {
			for _, customType := range types {
				if customType.IsPrimitive {
					result = append(result, ReplaceGrizzlyId(bytes, customType.Name)...)
				}
			}
		} else {
			result = append(result, bytes...)
		}
	}

	result = RemovePackage(result)
	result = ReplaceImports(result)

	return result, err
}

func GenCollectionCode(config GrizzlyConfigCollection, isSimple bool) (result string, err error) {
	code, err := GetCollectionCode()
	types := GenerateTypes(config.Types)

	if err != nil {
		return result, err
	}

	methodCode, err := GetMethodsCode(config.Methods, types)

	code = append(code, methodCode...)
	code = InjectImports(code, GetImportsByMethods(config.Methods))
	code = append([]byte("package " + config.Package), code...)

	code = GenCode(&config, code, isSimple)

	if err != nil {
		return result, err
	}

	code = ReplaceSearchCallback(code, config.Name)

	return string(code), err
}

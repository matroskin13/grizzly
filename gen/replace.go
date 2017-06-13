package gen

import (
	"regexp"
	"strings"
)

func ReplaceModel(code []byte, modelName string, types map[string]string) []byte {
	var structString = " {\n"

	for key, value := range types {
		structString += "\t" + strings.Title(key) + " " + value + "\n"
	}

	structString += "}";

	r, _ := regexp.Compile("type Model struct \\{GrizzlyId int; GrizzlyName string}")
	result := r.ReplaceAll(code, []byte("type Model struct" + structString))

	rModel, _ := regexp.Compile("Model")
	result = rModel.ReplaceAll(result, []byte(strings.Title(modelName)))

	return result
}

func ReplaceSearchCallback(code []byte, modelName string) []byte {
	rSearchCallback, _ := regexp.Compile("SearchCallback")
	result := rSearchCallback.ReplaceAll(code, []byte("SearchCallback" + strings.Title(modelName)))

	return result
}

func ReplaceCollection(code []byte, modelName string) []byte {
	rCollections, _ := regexp.Compile("Collection")
	result := rCollections.ReplaceAll(code, []byte(strings.Title(modelName) + "Collection"))

	rPackage, _ := regexp.Compile("package collection")
	result = rPackage.ReplaceAll(result, make([]byte, 0))

	return result
}

func ReplaceGrizzlyId(code []byte, customType string) []byte {
	gICollections, _ := regexp.Compile("GrizzlyId")
	result := gICollections.ReplaceAll(code, []byte(strings.Title(customType)))

	return result
}

func ReplaceImports(code []byte) []byte {
	rICollections, _ := regexp.Compile("(import \"sort\")")
	result := rICollections.ReplaceAll(code, []byte(""))

	return result
}

func InjectImports(code []byte, imports []string) (result []byte) {
	result = code[:]

	for _, i := range imports {
		result = append([]byte("\nimport \"" + i + "\"\n"), code...)
	}

	return result
}

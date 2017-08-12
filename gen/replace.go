package gen

import (
	"regexp"
	"strings"
)

func ReplaceSearchCallback(code []byte, modelName string) []byte {
	rSearchCallback, _ := regexp.Compile("SearchCallback")
	result := rSearchCallback.ReplaceAll(code, []byte("SearchCallback"+strings.Title(modelName)))

	return result
}

// Remove package name
func RemovePackage(code []byte) []byte {
	rPackage, _ := regexp.Compile("package collection")
	result := rPackage.ReplaceAll(code, make([]byte, 0))

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
		result = append([]byte("\nimport \""+i+"\"\n"), code...)
	}

	return result
}

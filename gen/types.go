package gen

import "strings"

type GrizzlyType struct {
	Name        string
	Value       string
	IsPointer   bool
	IsPrimitive bool
}

var PrimitiveTypes = []string{
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"uintptr",
	"string",
	"bool",
	"float64",
	"float32",
}

func GenerateTypes(configTypes map[string]string) []GrizzlyType {
	var types []GrizzlyType

	for key, value := range configTypes {
		var customType = GrizzlyType{
			Name:  key,
			Value: value,
		}

		if strings.Contains(value, "*") {
			customType.IsPointer = true
			customType.Value = value[1:]
		}

		for _, primitive := range PrimitiveTypes {
			if primitive == value {
				customType.IsPrimitive = true
			}
		}

		types = append(types, customType)
	}

	return types
}

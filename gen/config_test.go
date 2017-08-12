package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigByCode(t *testing.T) {
	code := `
package test

import (
	"fmt"
)

//grizzly:generate
type User struct {
	Id int
	Email string
}

//grizzly:generate
type Car struct {
	Brand string
	Model string
}

func GoGo() {
	fmt.Println("start")
}
	`

	expectedConfig := &GrizzlyConfig{
		Collections: []GrizzlyConfigCollection{
			{
				Name: "User",
				Types: map[string]string{
					"Id": "int",
					"Email": "string",
				},
				Package: "test",
				Methods: GetDefaultMethods(),
			},

			{
				Name: "Car",
				Types: map[string]string{
					"Brand": "string",
					"Model": "string",
				},
				Package: "test",
				Methods: GetDefaultMethods(),
			},
		},
	}

	config, err := GetConfigByCode([]byte(code))

	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

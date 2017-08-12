package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_MapToInt(t *testing.T) {
	items := GetCollection().MapToInt(func(item *Model) int {
		return item.GrizzlyId
	})

	assert.Equal(t, items, []int{1, 2, 3, 4})
}

func TestCollection_MapToString(t *testing.T) {
	items := GetCollection().MapToString(func(item *Model) string {
		return item.GrizzlyName
	})

	assert.Equal(t, items, []string{"item1", "item2", "item3", "item4"})
}

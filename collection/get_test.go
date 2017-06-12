package collection

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCollection_Get(t *testing.T) {
	item := GetCollection().Get(1)

	assert.Equal(t, item.GrizzlyId, 2)

	item = GetCollection().Get(20)

	assert.Nil(t, item)
}

package collection

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCollection_Push(t *testing.T) {
	collection := GetCollection().Push(&Model{GrizzlyId: 10, GrizzlyName: "test"})

	assert.Len(t, collection.Items, 5)
	assert.Equal(t, collection.Get(4).GrizzlyId, 10)
}

func TestCollection_Pop(t *testing.T) {
	collection := GetCollection()
	item := collection.Pop()

	assert.Equal(t, item.GrizzlyId, 4)
	assert.Len(t, collection.Items, 3)
	assert.Equal(t, collection.Get(2).GrizzlyId, 3)
}

func TestCollection_Unshift(t *testing.T) {
	collection := GetCollection().Unshift(&Model{GrizzlyId: 10, GrizzlyName: "test"})

	assert.Len(t, collection.Items, 5)
	assert.Equal(t, collection.Get(0).GrizzlyId, 10)
}

func TestCollection_Shift(t *testing.T) {
	collection := GetCollection()
	item := collection.Shift()

	assert.Equal(t, item.GrizzlyId, 1)
	assert.Len(t, collection.Items, 3)
	assert.Equal(t, collection.Get(2).GrizzlyId, 4)
}

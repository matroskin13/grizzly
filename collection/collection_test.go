package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetCollection() *Collection {
	return NewCollection([]*Model{
		&Model{GrizzlyId: 1, GrizzlyName: "item1"},
		&Model{GrizzlyId: 2, GrizzlyName: "item2"},
		&Model{GrizzlyId: 3, GrizzlyName: "item3"},
		&Model{GrizzlyId: 4, GrizzlyName: "item4"},
	})
}

func TestNewCollection(t *testing.T) {
	var first = Model{GrizzlyId: 1, GrizzlyName: "first"}
	var second = Model{GrizzlyId: 2, GrizzlyName: "second"}

	expected := &Collection{Items: []*Model{&first, &second}}
	collection := NewCollection([]*Model{&first, &second})

	assert.Equal(t, expected, collection)
}

func TestNewEmptyCollection(t *testing.T) {
	collection := NewEmptyCollection()

	assert.Len(t, collection.Items, 0)
}

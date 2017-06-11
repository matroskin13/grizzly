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

func TestCollection_Filter(t *testing.T) {
	collection := GetCollection().Filter(func (item *Model) bool {
		return item.GrizzlyId > 2
	})

	assert.Len(t, collection.Items, 2)
	assert.Equal(t, collection.Items[0].GrizzlyId, 3)
	assert.Equal(t, collection.Items[1].GrizzlyId, 4)
}

func TestCollection_Find(t *testing.T) {
	item := GetCollection().Find(func (item *Model) bool {
		return item.GrizzlyId == 3
	})

	assert.Equal(t, item.GrizzlyId, 3)
}

func TestCollection_MapToInt(t *testing.T) {
	items := GetCollection().MapToInt(func (item *Model) int {
		return item.GrizzlyId
	})

	assert.Equal(t, items, []int{1, 2, 3, 4})
}

func TestCollection_MapToString(t *testing.T) {
	items := GetCollection().MapToString(func (item *Model) string {
		return item.GrizzlyName
	})

	assert.Equal(t, items, []string{"item1", "item2", "item3", "item4"})
}

func TestCollection_Get(t *testing.T) {
	item := GetCollection().Get(1)

	assert.Equal(t, item.GrizzlyId, 2)

	item = GetCollection().Get(20)

	assert.Nil(t, item)
}

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
package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_UniqByGrizzlyId(t *testing.T) {
	collection := NewCollection([]*Model{
		&Model{GrizzlyId: 1, GrizzlyName: "item1"},
		&Model{GrizzlyId: 2, GrizzlyName: "item2"},
		&Model{GrizzlyId: 2, GrizzlyName: "item3"},
		&Model{GrizzlyId: 4, GrizzlyName: "item4"},
	}).UniqByGrizzlyId()

	assert.Len(t, collection.Items, 3)
	assert.Equal(t, collection.Get(1).GrizzlyId, 2)
	assert.Equal(t, collection.Get(2).GrizzlyId, 4)
}

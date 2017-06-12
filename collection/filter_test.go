package collection

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCollection_Filter(t *testing.T) {
	collection := GetCollection().Filter(func (item *Model) bool {
		return item.GrizzlyId > 2
	})

	assert.Len(t, collection.Items, 2)
	assert.Equal(t, collection.Items[0].GrizzlyId, 3)
	assert.Equal(t, collection.Items[1].GrizzlyId, 4)
}

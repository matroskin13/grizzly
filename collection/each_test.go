package collection

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCollection_ForEach(t *testing.T) {
	var iterationCount int

	collection := GetCollection()

	collection.ForEach(func (item *Model) {
		assert.Equal(t, collection.Get(iterationCount).GrizzlyId, item.GrizzlyId)
		iterationCount += 1
	})

	assert.Equal(t, iterationCount, 4)
}

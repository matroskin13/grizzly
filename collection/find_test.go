package collection

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCollection_Find(t *testing.T) {
	item := GetCollection().Find(func (item *Model) bool {
		return item.GrizzlyId == 3
	})

	assert.Equal(t, item.GrizzlyId, 3)
}

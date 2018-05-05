package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_Sort(t *testing.T) {
	sortByIdAsc := GetCollection().SortByGrizzlyId("asc")
	sortByIdDesc := GetCollection().SortByGrizzlyId("desc")
	qsortedByIdAsc := GetCollection().Sort(func(a *Model, b *Model) int {
		if a.GrizzlyId > b.GrizzlyId {
			return 1
		} else if a.GrizzlyId == b.GrizzlyId {
			return 0
		} else {
			return -1
		}
	})
	qsortedByIdDesc := GetCollection().Sort(func(a *Model, b *Model) int {
		if a.GrizzlyId < b.GrizzlyId {
			return 1
		} else if a.GrizzlyId == b.GrizzlyId {
			return 0
		} else {
			return -1
		}
	})

	assert.Equal(t, qsortedByIdAsc, sortByIdAsc)
	assert.Equal(t, qsortedByIdDesc, sortByIdDesc)
}

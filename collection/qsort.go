package collection

func (c *Collection) Sort(comparator func(a *Model, b *Model) int) *Collection {
	collection := &Collection{Items: c.Items}
	qsort(comparator, collection.Items, 0, len(collection.Items)-1)
	return collection
}

// Quicksort
func qsort(comparator func(a *Model, b *Model) int, collection []*Model, lo int, hi int) {
	if lo < hi {
		p := partition(comparator, collection, lo, hi)
		qsort(comparator, collection, lo, p)
		qsort(comparator, collection, p+1, hi)
	}
}

// Hoare partition scheme
func partition(comparator func(a *Model, b *Model) int, collection []*Model, lo int, hi int) int {
	pivot := collection[lo]
	i := lo - 1
	j := hi + 1
	for {
		for ok := true; ok; ok = (comparator(collection[i], pivot) == -1) {
			i++
		}
		for ok := true; ok; ok = (comparator(collection[j], pivot) == 1) {
			j--
		}

		if i >= j {
			return j
		}
		collection[i], collection[j] = collection[j], collection[i]
	}
}

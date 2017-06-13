package collection

import "sort"

type byGrizzlyIdAsc []*Model

func (a byGrizzlyIdAsc) Len() int           { return len(a) }
func (a byGrizzlyIdAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byGrizzlyIdAsc) Less(i, j int) bool { return a[i].GrizzlyId < a[j].GrizzlyId }

type byGrizzlyIdDesc []*Model

func (a byGrizzlyIdDesc) Len() int           { return len(a) }
func (a byGrizzlyIdDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byGrizzlyIdDesc) Less(i, j int) bool { return a[i].GrizzlyId > a[j].GrizzlyId }

func (c *Collection) SortByGrizzlyId(mode string) *Collection {
	collection := NewCollection(c.Items)

	if mode == "desc" {
		sort.Sort(byGrizzlyIdDesc(collection.Items))
	} else {
		sort.Sort(byGrizzlyIdAsc(collection.Items))
	}

	return collection
}
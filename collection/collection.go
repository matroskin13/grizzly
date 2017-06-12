package collection

type Model struct {GrizzlyId int; GrizzlyName string}

type Collection struct {
	Items []*Model
}

type SearchCallback func(item *Model) bool

func NewCollection(items []*Model) *Collection {
	var collection Collection

	collection.Items = items

	return &collection
}

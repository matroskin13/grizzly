package collection

type Model struct {GrizzlyId int; GrizzlyName string}

type Collection struct {
	Items []*Model
}

type SearchCallback func(item *Model) bool

//grizzly:replaceName New{{.Name}}Collection
func NewCollection(items []*Model) *Collection {
	var collection Collection

	collection.Items = items

	return &collection
}

//grizzly:replaceName NewEmpty{{.Name}}Collection
func NewEmptyCollection() *Collection {
	return &Collection{}
}

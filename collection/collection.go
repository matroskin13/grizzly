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

func (c *Collection) Filter(callback SearchCallback) *Collection {
	var newItems []*Model

	for _, v := range c.Items {
		if callback(v) == true {
			newItems = append(newItems, v)
		}
	}

	return &Collection{Items: newItems}
}

func (c *Collection) Find(callback SearchCallback) *Model {
	for _, v := range c.Items {
		if callback(v) == true {
			return v
		}
	}

	return nil
}

func (c *Collection) MapToInt(callback func(item *Model) int) []int {
	items := []int{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *Collection) MapToString(callback func(item *Model) string) []string {
	items := []string{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}
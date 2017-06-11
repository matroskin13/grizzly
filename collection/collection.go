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

func (c *Collection) Get(index int) (model *Model) {
	if len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

func (c *Collection) Push(item *Model) *Collection {
	newItems := append(c.Items, item)

	return &Collection{Items: newItems}
}

func (c *Collection) Shift() *Model {
	item := c.Items[0];
	c.Items = c.Items[1:];

	return item
}

func (c *Collection) Pop() *Model {
	item := c.Items[len(c.Items) - 1];
	c.Items = c.Items[:len(c.Items) - 1];

	return item
}

func (c *Collection) Unshift(item *Model) *Collection {
	newItems := append([]*Model{item}, c.Items...)

	return &Collection{Items: newItems}
}
package collection

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

func (c *Collection) Len() int {
	return len(c.Items)
}
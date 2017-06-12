package collection

func (c *Collection) Filter(callback SearchCallback) *Collection {
	var newItems []*Model

	for _, v := range c.Items {
		if callback(v) == true {
			newItems = append(newItems, v)
		}
	}

	return &Collection{Items: newItems}
}

package collection

func (c *Collection) ForEach(callback func (item *Model)) {
	for _, i := range c.Items {
		callback(i)
	}
}

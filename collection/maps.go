package collection

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

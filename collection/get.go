package collection

func (c *Collection) Get(index int) (model *Model) {
	if len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

package collection

func (c *Collection) Get(index int) (model *Model) {
	if index >= 0 && len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

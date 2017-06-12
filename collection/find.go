package collection

func (c *Collection) Find(callback SearchCallback) *Model {
	for _, v := range c.Items {
		if callback(v) == true {
			return v
		}
	}

	return nil
}
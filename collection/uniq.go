package collection

func (c *Collection) UniqByGrizzlyId() *Collection {
	collection := NewCollection([]*Model{})

	for _, item := range c.Items {
		searchItem := collection.Find(func (model *Model) bool {
			return model.GrizzlyId == item.GrizzlyId
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

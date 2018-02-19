package main
import (
	"github.com/icrowley/fake"
  )

func createCategoryFactory(c Category) Category {
	name := fake.Industry()
	if(c.name != "") {
		name = c.name
	}
	data := Category{
		name: name,
	}
	id := InsertNewCategory(data)
	data.id = id

	return data
}
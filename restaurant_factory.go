package main
import (
	"github.com/icrowley/fake"
  )

func createRestaurantFactory(r Restaurant) Restaurant {
	rating := 2
	name := fake.Company()
	longitude := float64(fake.Longitude())
	latitude := float64(fake.Latitude())
	idCategory := r.Id_category

	if(r.Id_category == 0){
		category := createCategoryFactory(Category{})
		idCategory = category.id
	}

	if(r.Rating != 0) {
		rating = r.Rating
	}
	if(r.Name != "") {
		name = r.Name
	}
	if(r.Longitude != 0){
		longitude = r.Longitude
	}
	if(r.Latitude != 0){
		latitude = r.Latitude
	}
	data := Restaurant{
		Name: name,
		Rating: rating,
		Longitude: longitude,
		Latitude: latitude,
		Id_category: idCategory,
	}
	id := InsertNewRestaurant(data)
	data.Id = id

	return data
}
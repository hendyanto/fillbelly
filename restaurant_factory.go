package main
import (
	"github.com/icrowley/fake"
  )

func createRestaurantFactory(r Restaurant) Restaurant {
	rating := 2
	name := fake.Company()
	if(r.Rating != 0) {
		rating = r.Rating
	}
	if(r.Name != "") {
		name = r.Name
	}
	data := Restaurant{
		Name: name,
		Rating: rating,
		Longitude: r.Longitude,
		Latitude: r.Latitude,
		Id_category: r.Id_category,
	}
	id := InsertNewRestaurant(data)
	data.Id = id

	return data
}
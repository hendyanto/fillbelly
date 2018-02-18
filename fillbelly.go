package main

import (
	"fmt"
)

func main() {
	InsertNewRestaurant(Restaurant{
		name: "Troll",
		id_category: 1,
		longitude: "6.9",
		latitude: "2.2",
		rating: 9,
	})
	results := GetAllRestaurant()
	for _, element := range results {
		fmt.Printf("%s | %d | %s | %s | %d \n", element.name, element.rating, element.longitude, element.latitude, element.id_category)
	}
}
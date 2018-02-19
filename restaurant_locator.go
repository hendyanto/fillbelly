package main

import (
	"math"
)

func getNearbyRestaurants(latitude float64, longitude float64) []Restaurant {
	all := GetAllRestaurant();
	var results []Restaurant

	userPosition := Position{
		latitude: latitude,
		longitude: longitude,
	}

	for _, element := range all {
		restaurantPosition := Position {
			latitude: element.Latitude,
			longitude: element.Longitude,
		}

		distance := calculateDistance(userPosition, restaurantPosition)
		if(distance < 5000) {
			results = append(results, element)
		}
	}
	return results
}

type Position struct {
	latitude float64
	longitude float64
}

func calculateDistance(pointA Position, pointB Position) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = pointA.latitude * math.Pi / 180
	lo1 = pointA.longitude * math.Pi / 180
	la2 = pointB.latitude * math.Pi / 180
	lo2 = pointB.longitude * math.Pi / 180
	r = 6378100

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	return 2 * r * math.Asin(math.Sqrt(h))
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
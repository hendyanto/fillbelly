package main

import (
	"time"
)

func reserve(name string, idRestaurant int, date time.Time) {
	InsertNewReservation(Reservation{
		name: name,
		id_restaurant: idRestaurant,
		date: date,
		created: time.Now(),
	})
}
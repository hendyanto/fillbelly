package main
import (
	"github.com/icrowley/fake"
	"time"
  )

func createReservationFactory(r Reservation) Reservation {
	name := fake.FullName()
	date := time.Now()
	created := time.Now()

	if(r.name != "") {
		name = r.name
	}
	beginningOfTime := time.Date(1,1,1,0,0,0,0,time.UTC)
	if(r.date != beginningOfTime) {
		date = r.date
	}
	if(r.created != beginningOfTime) {
		created = r.created
	}
	data := Reservation{
		name: name,
		id_restaurant: r.id_restaurant,
		date: date,
		created: created,
	}
	id := InsertNewReservation(data)
	data.id = id

	return data
}
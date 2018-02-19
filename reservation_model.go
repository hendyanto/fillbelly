package main

import (
	"fmt"
	"time"
	"database/sql"
	pq "github.com/lib/pq"
)

type Reservation struct {
    id int
    name string
	id_restaurant int
	date time.Time
	created time.Time
}

func InsertNewReservation(r Reservation) int {
	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("reservations")
	var id int
	err := db.QueryRow(fmt.Sprintf("INSERT INTO %s (name, id_restaurant, date, created) VALUES ($1, $2, $3, $4) RETURNING id", quoted), r.name, r.id_restaurant, r.date, r.created).Scan(&id)
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	return id
}
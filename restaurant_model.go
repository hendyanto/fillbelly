package main

import (
	"fmt"
	"database/sql"
	pq "github.com/lib/pq"
)

type Restaurant struct {
    id int
    name string
    rating int
    id_category int
    longitude string
    latitude string
}

func InsertNewRestaurant(r Restaurant){
	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("restaurants")
	_, err := db.Exec(fmt.Sprintf("INSERT INTO %s (name,rating,longitude,latitude,id_category) VALUES ($1, $2, $3, $4, $5)", quoted), r.name, r.rating, r.longitude, r.latitude, r.id_category)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func GetAllRestaurant() []Restaurant {

	var results []Restaurant

	db, err := sql.Open("postgres", DbConnectionString());
	if err != nil {
		fmt.Println("pq error: ?", err)
	} else {
		rows, err_sql := db.Query("SELECT restaurants.* FROM restaurants INNER JOIN categories on restaurants.id_category = categories.id;")

		if err_sql != nil {
			fmt.Println("Query Error: ?", err_sql)
		} else {
			for rows.Next() {
				r := Restaurant {}

				err = rows.Scan(&r.id, &r.name, &r.rating, &r.longitude, &r.latitude, &r.id_category)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				results = append(results, Restaurant {
					name: r.name,
					rating: r.rating,
					longitude: r.longitude,
					latitude: r.latitude,
					id_category: r.id_category })
			}
		}
	}

	return results
}
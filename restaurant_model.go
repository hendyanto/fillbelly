package main

import (
	"fmt"
	"database/sql"
	pq "github.com/lib/pq"
)

type Restaurant struct {
    Id int
    Name string
    Rating int
    Id_category int
    Longitude float64
    Latitude float64
    Category_name string
}

func InsertNewRestaurant(r Restaurant) int {
	var id int
	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("restaurants")
	err := db.QueryRow(fmt.Sprintf("INSERT INTO %s (name,rating,longitude,latitude,id_category) VALUES ($1, $2, $3, $4, $5) RETURNING id", quoted), r.Name, r.Rating, r.Longitude, r.Latitude, r.Id_category).Scan(&id)
	db.Close()
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	return id
}

func GetAllRestaurant() []Restaurant {

	var results []Restaurant

	db, err := sql.Open("postgres", DbConnectionString());
	if err != nil {
		fmt.Println("pq error: ?", err)
	} else {
		rows, err_sql := db.Query("SELECT restaurants.*, categories.name AS category_name FROM restaurants INNER JOIN categories on restaurants.id_category = categories.id;")
		db.Close()

		if err_sql != nil {
			fmt.Println("Query Error: ?", err_sql)
		} else {
			for rows.Next() {
				r := Restaurant {}

				err = rows.Scan(&r.Id, &r.Name, &r.Rating, &r.Longitude, &r.Latitude, &r.Id_category, &r.Category_name)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				results = append(results, Restaurant {
					Id: r.Id,
					Name: r.Name,
					Rating: r.Rating,
					Longitude: r.Longitude,
					Latitude: r.Latitude,
					Id_category: r.Id_category,
					Category_name: r.Category_name,
				 })
			}
		}
	}

	return results
}
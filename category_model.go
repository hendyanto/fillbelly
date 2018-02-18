package main

import (
	"fmt"
	"database/sql"
	pq "github.com/lib/pq"
)

type Category struct {
    id int
    name string
}

func InsertNewCategory(r Category) int {
	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("categories")
	var id int
	err := db.QueryRow(fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", quoted), r.name).Scan(&id)
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	return id
}
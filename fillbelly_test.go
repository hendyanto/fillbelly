package main 

import (
	"fmt"
	"testing"
	"database/sql"
	"os"
	pq "github.com/lib/pq"
)

func isEnvTest() bool{
	return (os.Getenv("GO_ENV") == "test")
}

func ResetDB() {
	if(isEnvTest()) {
		clearRestaurantTable()
		// createRestaurantTable()
		// createCategoryTable()
	} else {
		fmt.Printf("ENV is not test, aborting db reset");
	}
}

func clearRestaurantTable(){
	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("restaurants")
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", quoted))
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func createRestaurantTable(){
	db, _ := sql.Open("postgres", DbConnectionString());
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE restaurants ( id serial primary key, name varchar(200), rating integer, longitude varchar(50), latitude varchar(50), id_category int);"))
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func createCategoryTable(){
	db, _ := sql.Open("postgres", DbConnectionString());
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE categories ( id serial primary key, name varchar(200));"))
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func TestRestaurantModel(t *testing.T) {
	if(!isEnvTest()) { return }
	ResetDB()

	idCategory := InsertNewCategory(Category{ name: "Seafood" })

	restaurantName := "Puvlic"
	restaurantRating := 9
	restaurantLongitude := "6.998"
	restaurantLatitude := "688.998"
	restaurantCategory := idCategory

	InsertNewRestaurant(Restaurant{
		name: restaurantName,
		rating: restaurantRating,
		longitude: restaurantLongitude,
		latitude: restaurantLatitude,
		id_category: restaurantCategory,
	})

	all := GetAllRestaurant()

	first := all[0]

	if( first.name != restaurantName ||
		first.id_category != restaurantCategory ||
		first.rating != restaurantRating ||
		first.longitude != restaurantLongitude ||
		first.latitude != restaurantLatitude) {
		t.Errorf("Some element did not match")
	}
}


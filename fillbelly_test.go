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
		dropRestaurantTable()
		dropCategoryTable()
		createRestaurantTable()
		createCategoryTable()
	} else {
		fmt.Printf("ENV is not test, aborting db reset");
	}
}

func dropRestaurantTable(){
	if(!isEnvTest()) { return }

	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("restaurants")
	_, err := db.Exec(fmt.Sprintf("DROP TABLE %s", quoted))
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func dropCategoryTable(){
	if(!isEnvTest()) { return }

	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("categories")
	_, err := db.Exec(fmt.Sprintf("DROP TABLE %s", quoted))
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func clearRestaurantTable(){
	if(!isEnvTest()) { return }
	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("restaurants")
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", quoted))
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}

func createRestaurantTable(){
	db, _ := sql.Open("postgres", DbConnectionString());
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE restaurants ( id serial primary key, name varchar(200), rating integer, longitude float, latitude float, id_category int);"))
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

func TestRestaurantLocator(t *testing.T) {
	if(!isEnvTest()) { return }
	ResetDB()

	categoryName := "Seafood"
	idCategory := InsertNewCategory(Category{ name: categoryName })
	
	restaurantName := "Puvlic"
	restaurantRating := 9
	restaurantLatitude := -6.115894
	restaurantLongitude := 106.7867153
	restaurantCategory := idCategory

	InsertNewRestaurant(Restaurant{
		Name: restaurantName,
		Rating: restaurantRating,
		Longitude: restaurantLongitude,
		Latitude: restaurantLatitude,
		Id_category: restaurantCategory,
	})

	farRestaurantName := "Masda"
	farRestaurantRating := 9
	farRestaurantLatitude := -6.174722
	farRestaurantLongitude := 106.7903383
	farRestaurantCategory := idCategory

	InsertNewRestaurant(Restaurant{
		Name: farRestaurantName,
		Rating: farRestaurantRating,
		Longitude: farRestaurantLongitude,
		Latitude: farRestaurantLatitude,
		Id_category: farRestaurantCategory,
	})

	latitude := -6.115734
	longitude := 106.7916073

	results := getNearbyRestaurants(latitude, longitude)
	var matched bool
	matched = false
	for _, element := range results {
		if(element.Name == farRestaurantName) {
			t.Errorf("Restaurant %s should not be included.", farRestaurantName);
		}
	}

	for _, element := range results {
		if(element.Name == restaurantName) {
			matched = true
		}

	}
	if !matched {
		t.Errorf("Restaurant %s should be included.", restaurantName);
	}
}

func TestRestaurantModel(t *testing.T) {
	if(!isEnvTest()) { return }
	ResetDB()

	categoryName := "Seafood"

	idCategory := InsertNewCategory(Category{ name: categoryName })

	restaurantName := "Puvlic"
	restaurantRating := 9
	restaurantLongitude := 6.998
	restaurantLatitude := 688.998
	restaurantCategory := idCategory

	InsertNewRestaurant(Restaurant{
		Name: restaurantName,
		Rating: restaurantRating,
		Longitude: restaurantLongitude,
		Latitude: restaurantLatitude,
		Id_category: restaurantCategory,
	})

	all := GetAllRestaurant()

	first := all[0]

	if( first.Name != restaurantName ||
		first.Id_category != restaurantCategory ||
		first.Rating != restaurantRating ||
		first.Longitude != restaurantLongitude ||
		first.Category_name != categoryName ||
		first.Latitude != restaurantLatitude) {
		t.Errorf("Some element did not match")
	}
}


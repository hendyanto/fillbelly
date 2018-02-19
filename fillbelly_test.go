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
		dropReservationTable()
		dropCategoryTable()
		createRestaurantTable()
		createCategoryTable()
		createReservationTable()
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

func dropReservationTable(){
	if(!isEnvTest()) { return }

	db, _ := sql.Open("postgres", DbConnectionString());
	quoted := pq.QuoteIdentifier("reservations")
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

func createReservationTable(){
	db, _ := sql.Open("postgres", DbConnectionString());
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE reservations ( id serial primary key, name varchar(200), id_restaurant integer, date timestamp, created timestamp);"))
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

	category := createCategoryFactory(Category{ })
	
	restaurantLatitude := -6.115894
	restaurantLongitude := 106.7867153

	restaurant := createRestaurantFactory(Restaurant{
		Longitude: restaurantLongitude,
		Latitude: restaurantLatitude,
		Id_category: category.id,
	})

	farRestaurantLatitude := -6.174722
	farRestaurantLongitude := 106.7903383

	farRestaurant := createRestaurantFactory(Restaurant{
		Longitude: farRestaurantLongitude,
		Latitude: farRestaurantLatitude,
		Id_category: category.id,
	})

	latitude := -6.115734
	longitude := 106.7916073

	results := getNearbyRestaurants(latitude, longitude)
	var matched bool
	matched = false
	for _, element := range results {
		if(element.Name == farRestaurant.Name) {
			t.Errorf("Restaurant %s should not be included.", farRestaurant.Name);
		}
	}

	for _, element := range results {
		if(element.Name == restaurant.Name) {
			matched = true
		}

	}
	if !matched {
		t.Errorf("Restaurant %s should be included.", restaurant.Name);
	}
}

func TestRestaurantModel(t *testing.T) {
	if(!isEnvTest()) { return }
	ResetDB()

	category := createCategoryFactory(Category{ })

	restaurantLongitude := 6.998
	restaurantLatitude := 688.998

	restaurant := createRestaurantFactory(Restaurant{
		Longitude: restaurantLongitude,
		Latitude: restaurantLatitude,
		Id_category: category.id,
	})

	all := GetAllRestaurant()

	first := all[0]

	if( first.Name != restaurant.Name ||
		first.Id_category != restaurant.Id_category ||
		first.Rating != restaurant.Rating ||
		first.Longitude != restaurantLongitude ||
		first.Category_name != category.name ||
		first.Latitude != restaurantLatitude) {
		t.Errorf("Some element did not match")
	}
}

func TestReservationModel(t *testing.T) {
	if(!isEnvTest()) { return }
	ResetDB()

	category := createCategoryFactory(Category{ })
	restaurant := createRestaurantFactory(Restaurant{
		Id_category: category.id,
	})

	reservation := createReservationFactory(Reservation{
		id_restaurant: restaurant.Id,
	})
	
	if(reservation.id == 0) {
		t.Errorf("Reservation is not created.")
	}
}
package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net"
	"strconv"
	"net/http"
	"time"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello!\r\n"))
	})

	sm.HandleFunc("/nearby", func(w http.ResponseWriter, r *http.Request) {
		latitude, latitudePresent := r.URL.Query()["latitude"]
		longitude, longitudePresent := r.URL.Query()["longitude"]
		if latitudePresent && longitudePresent {
			w.WriteHeader(200)
			fmt.Printf("\nLatitude: %s\nLongitude: %s\n", latitude[0], longitude[0])
			w.Write(getNearby(latitude[0], longitude[0]))
		}
	})

	sm.HandleFunc("/reserve", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			name, namePresent := r.Form["name"]
			idRestaurant, idPresent := r.Form["id_restaurant"]
			date, datePresent := r.Form["date"]

			if namePresent && idPresent && datePresent {
				parsedDate, _ := time.Parse("2006-01-02T15:04:05Z", date[0])
				w.WriteHeader(200)
				converted, _ := strconv.Atoi(idRestaurant[0])
				reserve(name[0], converted, parsedDate)
			} else {
				w.WriteHeader(400)
			}
		}
	})
	
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, sm))
}

func getNearby(latitude string, longitude string) []byte {
	latitudeFloat, _ := strconv.ParseFloat(latitude, 64)
	longitudeFloat, _ := strconv.ParseFloat(longitude, 64)

	b, _ := json.Marshal(getNearbyRestaurants(latitudeFloat, longitudeFloat))
	return b
}
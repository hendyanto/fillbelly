package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net"
	"strconv"
	"net/http"
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
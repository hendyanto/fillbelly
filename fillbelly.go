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
	fmt.Print("Starting HTTP Server\n");
	fmt.Print(`:::::::::: ::::::::::: :::        :::        :::::::::  :::::::::: :::        :::        :::   ::: 
:+:            :+:     :+:        :+:        :+:    :+: :+:        :+:        :+:        :+:   :+: 
+:+            +:+     +:+        +:+        +:+    +:+ +:+        +:+        +:+         +:+ +:+  
:#::+::#       +#+     +#+        +#+        +#++:++#+  +#++:++#   +#+        +#+          +#++:   
+#+            +#+     +#+        +#+        +#+    +#+ +#+        +#+        +#+           +#+    
#+#            #+#     #+#        #+#        #+#    #+# #+#        #+#        #+#           #+#    
###        ########### ########## ########## #########  ########## ########## ##########    ###    `)
	fmt.Print("\n")
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello!\r\n"))
	})

	sm.HandleFunc("/nearby", NearbyHandler)

	sm.HandleFunc("/reserve", ReserveHandler)
	
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, sm))
}

func ReserveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n POST /reserve")
	if r.Method == "POST" {
		r.ParseForm()
		name, namePresent := r.Form["name"]
		idRestaurant, idPresent := r.Form["id_restaurant"]
		date, datePresent := r.Form["date"]

		if namePresent && idPresent && datePresent && name[0] != "" && idRestaurant[0] != "" && date[0] != "" {
			parsedDate, _ := time.Parse("2006-01-02T15:04:05Z", date[0])
			w.WriteHeader(200)
			converted, _ := strconv.Atoi(idRestaurant[0])
			reserve(name[0], converted, parsedDate)
		} else {
			w.WriteHeader(400)
		}
	}
}

func NearbyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n GET /nearby")
	latitude, latitudePresent := r.URL.Query()["latitude"]
	longitude, longitudePresent := r.URL.Query()["longitude"]
	if latitudePresent && longitudePresent {
		w.WriteHeader(200)
		w.Write(getNearby(latitude[0], longitude[0]))
	} else {
		w.WriteHeader(400)
	}
}

func getNearby(latitude string, longitude string) []byte {
	latitudeFloat, _ := strconv.ParseFloat(latitude, 64)
	longitudeFloat, _ := strconv.ParseFloat(longitude, 64)

	b, _ := json.Marshal(getNearbyRestaurants(latitudeFloat, longitudeFloat))
	return b
}
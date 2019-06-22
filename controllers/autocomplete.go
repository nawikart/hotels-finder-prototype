package controllers

import (
	"../db/psql"
	// "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func AutocompleteApi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	
	vars := mux.Vars(r)
	keyword := vars["keyword"]

	db, _ := psql.Connect()
	defer db.Close()

	var h Hotel
	var c City
	var a Autocomplete
	var as []Autocomplete

	//// cities
	rows, err := db.Query("SELECT * FROM cities WHERE city LIKE '%"+ keyword +"%' OR country LIKE '%"+ keyword +"%' LIMIT 5")
	if err != nil {
		panic(err)
	}	
	for rows.Next() {
		err = rows.Scan(&c.City_id, &c.City, &c.City_key, &c.Country, &c.CountryIsoCode, &c.Country_key, &c.HotelsCount)
		if err != nil {
			panic(err)
		}
		a.Type = "c"
		a.Value = c.City
		a.Key = c.City_key
		a.HotelsCount = c.HotelsCount

		as = append(as, a)
	}	
	
	/////// hotels
	rows, err = db.Query("SELECT * FROM hotels4listing WHERE hotel_name LIKE '%"+ keyword +"%' ORDER BY rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC LIMIT 8")
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		err = rows.Scan(&h.Hotel_id, &h.Hotel_name, &h.City, &h.Country, &h.Countryisocode, &h.Hotel_name_key, &h.City_key, &h.Photo1, &h.Rates_from, &h.Rating_average,
			&h.Star_rating, &h.Number_of_reviews)
		if err != nil {
			panic(err)
		}
		a.Type = "h"
		a.Value = h.Hotel_name
		a.Key = h.Hotel_name_key
		a.Rating_average = h.Rating_average

		as = append(as, a)
	}

	js, _ := json.Marshal(as)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}
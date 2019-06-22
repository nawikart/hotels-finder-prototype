package controllers

import (
	"../db/psql"
	"regexp"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Listing struct {
	Ps, P int64
	Hotels []Hotel
	City, Country, Citykey, Countryisocode interface{}
}

func HotelsPerCity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	url := strings.ToLower(r.URL.String())
	rg, _ := regexp.Compile("/([-a-z]+)-([a-z][a-z])-hotels((-p([0-9]+))*).html")
	urlKeys := rg.FindStringSubmatch(url)
	// fmt.Println(urlKeys)
	citykey := urlKeys[1]
	countryIsoCode := urlKeys[2]

	// var page string
	// if urlKeys[5] != "" {
	// 	page = urlKeys[5]
	// }else{
	// 	page = "1"
	// }

	db, _ := psql.Connect()
	defer db.Close()

	var h Hotel
	var hs []Hotel
	var l Listing

	rows, err := db.Query("SELECT * FROM hotels4listing WHERE city_key = '"+ citykey +"' ORDER BY rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC LIMIT 24")
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		err = rows.Scan(&h.Hotel_id, &h.Hotel_name, &h.City, &h.Country, &h.Countryisocode, &h.Hotel_name_key, &h.City_key, &h.Photo1, &h.Rates_from, &h.Rating_average,
			&h.Star_rating, &h.Number_of_reviews)
		if err != nil {
			panic(err)
		}

		hs = append(hs, h)
	}

	l.P = 1
	l.Ps = 20
	l.Hotels = hs
	l.Countryisocode = countryIsoCode
	l.Citykey = citykey
	l.City = h.City
	l.Country = h.Country
	
	js, _ := json.Marshal(l)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

func HotelApi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	vars := mux.Vars(r)
	hotelnamekey := vars["hotelnamekey"]
	// city_cc := vars["city_cc"]

	db, _ := psql.Connect()
	defer db.Close()

	var h HotelDetail

	rows, err := db.Query(`SELECT hotel_id, hotel_name, city, country, countryisocode, hotel_name_key, city_key, rates_from, rating_average,
	star_rating, number_of_reviews, overview, photo1, photo2, photo3, photo4 FROM hotels WHERE hotel_name_key = '`+ hotelnamekey +`' LIMIT 1`)
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		err = rows.Scan(&h.Hotel_id, &h.Hotel_name, &h.City, &h.Country, &h.Countryisocode, &h.Hotel_name_key, &h.City_key, &h.Rates_from, &h.Rating_average,
			&h.Star_rating, &h.Number_of_reviews, &h.Overview, &h.Photo1, &h.Photo2, &h.Photo3, &h.Photo4)
		if err != nil {
			panic(err)
		}
	}
	
	js, _ := json.Marshal(h)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}
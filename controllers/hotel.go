package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nawikart/hotels-finder-prototype/db/psql"
)

type Listing struct {
	Ps, P                                  int64
	Hotels                                 []Hotel
	City, Country, Citykey, Countryisocode interface{}
}

func HotelsPerCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	url := strings.ToLower(r.URL.String())
	rg, _ := regexp.Compile("/([-a-z]+)-([a-z][a-z])-hotels((-p([0-9]+))*).html")
	urlKeys := rg.FindStringSubmatch(url)

	if len(urlKeys) > 0 {

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

		rows, err := db.Query("SELECT * FROM hotels4listing WHERE city_key = '" + citykey + "' ORDER BY rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC LIMIT 24")
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
}

func HotelsFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	vars := mux.Vars(r)
	citycountry := vars["citycountry"]
	datafilter := vars["datafilter"]

	url := strings.ToLower(citycountry)
	rg, _ := regexp.Compile("([-a-z]+)-([a-z][a-z])-hotels((-p([0-9]+))*).html")
	urlKeys := rg.FindStringSubmatch(url)

	citykey := urlKeys[1]

	db, _ := psql.Connect()
	defer db.Close()

	var h Hotel
	var hs []Hotel
	var l Listing

	//////////////
	datafilter_byt := []byte(datafilter)
	var datafilter_map map[string]string

	if err := json.Unmarshal(datafilter_byt, &datafilter_map); err != nil {
		panic(err)
	}

	var sql_dailyRate_min, sql_dailyRate_max, sql_minReviewScore, sql_minStarRating string
	sql_sortby := " ORDER BY rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC"
	if datafilter_map["sortBy"] != "" {
		switch datafilter_map["sortBy"] {
		case "PriceDesc":
			sql_sortby = " ORDER BY rates_from DESC, rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC"
		case "PriceAsc":
			sql_sortby = " ORDER BY rates_from ASC, rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC"
		case "StarRatingDesc":
			sql_sortby = " ORDER BY star_rating::float DESC, rating_average::float DESC, number_of_reviews DESC"
		case "StarRatingAsc":
			sql_sortby = " ORDER BY star_rating::float ASC, rating_average::float DESC, number_of_reviews DESC"
		case "ReviewScoreDesc":
			sql_sortby = " ORDER BY rating_average::float DESC, star_rating::float DESC, number_of_reviews DESC"
		case "ReviewScoreAsc":
			sql_sortby = " ORDER BY rating_average::float ASC, star_rating::float DESC, number_of_reviews DESC"
		case "ReviewsCountDesc":
			sql_sortby = " ORDER BY number_of_reviews DESC, rating_average::float DESC, star_rating::float DESC"
		case "ReviewsCountAsc":
			sql_sortby = " ORDER BY number_of_reviews ASC, rating_average::float DESC, star_rating::float DESC"
		}
	}
	if datafilter_map["dailyRate_min"] != "" {
		dailyRate_min := datafilter_map["dailyRate_min"]
		sql_dailyRate_min = " AND rates_from >= " + dailyRate_min
	}
	if datafilter_map["dailyRate_max"] != "" {
		dailyRate_max := datafilter_map["dailyRate_max"]
		sql_dailyRate_max = " AND (rates_from < " + dailyRate_max + " OR rates_from = " + dailyRate_max + ")"
	}
	if datafilter_map["minReviewScore"] != "" {
		minReviewScore := datafilter_map["minReviewScore"]
		sql_minReviewScore = " AND rating_average::float >= " + minReviewScore
	}
	if datafilter_map["minStarRating"] != "" {
		minStarRating := datafilter_map["minStarRating"]
		sql_minStarRating = " AND star_rating::float > " + minStarRating
	}

	rows, err := db.Query("SELECT * FROM hotels4listing WHERE city_key = '" + citykey + "'" + sql_dailyRate_min + sql_dailyRate_max + sql_minReviewScore + sql_minStarRating + sql_sortby + " LIMIT 24")
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
	l.Countryisocode = h.Countryisocode
	l.Citykey = h.City_key
	l.City = h.City
	l.Country = h.Country

	js, _ := json.Marshal(l)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HotelByNamekey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	vars := mux.Vars(r)
	hotelnamekey := vars["hotelnamekey"]
	// city_cc := vars["city_cc"]

	db, _ := psql.Connect()
	defer db.Close()

	var h HotelDetail

	rows, err := db.Query(`SELECT hotel_id, hotel_name, city, country, countryisocode, hotel_name_key, city_key, rates_from, rating_average,
	star_rating, number_of_reviews, overview, photo1, photo2, photo3, photo4 FROM hotels WHERE hotel_name_key = '` + hotelnamekey + `' LIMIT 1`)
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

func HotelByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	vars := mux.Vars(r)
	hotel_id := vars["hotel_id"]

	db, _ := psql.Connect()
	defer db.Close()

	var h HotelDetail

	rows, err := db.Query(`SELECT hotel_id, hotel_name, city, country, countryisocode, hotel_name_key, city_key, rates_from, rating_average,
	star_rating, number_of_reviews, overview, photo1, photo2, photo3, photo4 FROM hotels WHERE hotel_id = '` + hotel_id + `' LIMIT 1`)
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

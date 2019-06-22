package controllers

import (
	"../db/psql"
	// "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Top24Cities(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	db, _ := psql.Connect()
	defer db.Close()

	var c City
	var cs []City

	rows, err := db.Query("SELECT * FROM top24cities LIMIT 12")
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		err = rows.Scan(&c.City_id, &c.City, &c.City_key, &c.Country, &c.CountryIsoCode, &c.Country_key, &c.HotelsCount)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)
	}
	
	js, _ := json.Marshal(cs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

func Top24CitiesPerCountry(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	
	vars := mux.Vars(r)
	countrykey := vars["countrykey"]
	// fmt.Println(countrykey)

	db, _ := psql.Connect()
	defer db.Close()

	var c City
	var cs []City

	rows, err := db.Query("SELECT * FROM cities WHERE country_key = '"+ countrykey +"' ORDER BY hotels_count DESC LIMIT 12")
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		err = rows.Scan(&c.City_id, &c.City, &c.City_key, &c.Country, &c.CountryIsoCode, &c.Country_key, &c.HotelsCount)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)
	}
	
	js, _ := json.Marshal(cs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

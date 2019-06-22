package controllers

import (
	"../db/psql"
	"net/http"
	"encoding/json"
)

func Top24Countries(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	db, _ := psql.Connect()
	defer db.Close()

	var c Country
	var cs []Country

	rows, err := db.Query("SELECT * FROM top24countries LIMIT 12")
	if err != nil {
		panic(err)
	}
	
	for rows.Next() {
		err = rows.Scan(&c.Country_id, &c.Country, &c.Countryisocode, &c.Country_key, &c.HotelsCount)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)
	}
	
	js, _ := json.Marshal(cs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}
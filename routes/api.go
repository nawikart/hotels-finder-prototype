package routes

import (
	"github.com/gorilla/mux"
	ctx "github.com/nawikart/hotels-finder-prototype/controllers"
)

func APIs(r *mux.Router) {

	r.HandleFunc("/api/autocomplete/{keyword}", ctx.AutocompleteApi).Methods("GET")
	r.HandleFunc("/api/getTop24Countries", ctx.Top24Countries).Methods("GET")
	r.HandleFunc("/api/getTop24Cities", ctx.Top24Cities).Methods("GET")
	r.HandleFunc("/api/getTop24CitiesPerCountry/{countrykey}", ctx.Top24CitiesPerCountry).Methods("GET")
	r.HandleFunc("/api/hotels-per-city/{citykey}", ctx.HotelsPerCity).Methods("GET")
	r.HandleFunc("/api/hotels/{citycountry}/filter/{datafilter}", ctx.HotelsFilter).Methods("GET")
	r.HandleFunc("/api/hotel/{hotelnamekey}/{city_cc}", ctx.HotelApi).Methods("GET")
}

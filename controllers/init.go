package controllers


type City struct {
	City_id, City, City_key, Country, CountryIsoCode, Country_key, HotelsCount interface{}
}

type Country struct {
	Country_id, Country, Countryisocode, Country_key, HotelsCount string
}

type Hotel struct{
	Hotel_id, Hotel_name, City, Country, Countryisocode, Hotel_name_key, City_key, Photo1, Rates_from, Rating_average,
	Star_rating, Number_of_reviews interface {}
}

type HotelDetail struct{
	Hotel_id, Hotel_name, City, Country, Countryisocode, Hotel_name_key, City_key, Rates_from, Rating_average,
	Star_rating, Number_of_reviews, Overview, Photo1, Photo2, Photo3, Photo4 interface {}
}

type Autocomplete struct{
	Key, Type, Value, Count, Rating_average, HotelsCount interface {}
}
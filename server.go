package main

import (
        "fmt"
        "github.com/gorilla/mux"
        "github.com/nawikart/hotels-finder-prototype/routes"
        "net/http"
		)
		    
func main() {
    r := mux.NewRouter()
    r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

    routes.APIs(r)
    routes.Nuxtjs(r)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", r)
}
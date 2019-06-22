package main

import (
        "fmt"
        "github.com/gorilla/mux"
        "./routes"
        "net/http"
		)
		    
func main() {
    r := mux.NewRouter()
    r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

    routes.APIs(r)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", r)
}
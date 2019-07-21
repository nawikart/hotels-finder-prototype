package routes

import (
	"net/http"
    "github.com/gorilla/mux"
    "html/template"
    "fmt"
)


func Nuxtjs(r *mux.Router) {
    
    r.PathPrefix("/_nuxt").Handler(http.StripPrefix("/_nuxt", http.FileServer(http.Dir("./nuxtjs/dist/_nuxt"))))

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/{one}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/{one}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/{one}/{two}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })                       
    r.HandleFunc("/{one}/{two}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
	})
	
    r.HandleFunc("/{one}/{two}/{three}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })                       
    r.HandleFunc("/{one}/{two}/{three}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })	 	
}
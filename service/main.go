package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muskankhedia/Revolutionize_Medicare/service/controllers"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func error404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `404 page not found`)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `Check`)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginHandler)
	r.HandleFunc("/profile", controllers.ProfileHandler)
	r.HandleFunc("/add_data", controllers.AddDataHandler)

	r.NotFoundHandler = http.HandlerFunc(error404)

	log.Fatal(http.ListenAndServe(":9000", r))
}

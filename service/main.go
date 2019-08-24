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
	r.HandleFunc("/get_profile", controllers.GetProfileHandler)
	r.HandleFunc("/update_profile", controllers.UpdateProfileHandler)
	r.HandleFunc("/add_data", controllers.AddDataHandler)
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/allevents", controllers.AllEventsHandler)

	r.NotFoundHandler = http.HandlerFunc(error404)

	log.Fatal(http.ListenAndServe(":9000", r))
}

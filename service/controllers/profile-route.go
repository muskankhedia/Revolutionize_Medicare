package controllers

import (
	"log"
	"net/http"
)

//ProfileHandler returns the profile details of the patient id
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In ProfileHandler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

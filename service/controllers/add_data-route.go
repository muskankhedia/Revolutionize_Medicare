package controllers

import (
	"fmt"
	"net/http"
)

//AddDataHandler adds the data into the user details 
func AddDataHandler(w http.ResponseWriter, r *http.Request) {

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	
	fmt.Println("Reached Add Data")
}
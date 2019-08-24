package main

import (
	"fmt"
	"log"
	"net/http"
	// "time"
	// "sync"

	"github.com/gorilla/mux"
	// "github.com/davecgh/go-spew/spew"
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

	// var mutex = &sync.Mutex{}
	// go func() {
	// 	t := time.Now()
	// 	genesisBlock := controllers.Block{}
	// 	genesisBlock = controllers.Block{0, t.String(), 0, controllers.CalculateHash(genesisBlock), "", "", []string{}, 0}
	// 	spew.Dump(genesisBlock)

	// 	mutex.Lock()
	// 	controllers.Blockchain = append(controllers.Blockchain, genesisBlock)
	// 	mutex.Unlock()
	// }()
	
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginHandler)
	r.HandleFunc("/profile", controllers.ProfileHandler)
	r.HandleFunc("/add_data", controllers.AddDataHandler)

	r.NotFoundHandler = http.HandlerFunc(error404)

	log.Fatal(http.ListenAndServe(":9000", r))
}

/*
package main


import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Ping struct {
	Value string `json:"value"`
}

var xRegCount int = 0
var ping = Ping{}

func main() {
	fmt.Println("Pinggg")
	ping = Ping{"ping"}

	r := mux.NewRouter()

	// seeding

	r.HandleFunc("/ping", increaseRegCount).Methods("GET")
	r.HandleFunc("/ping", show405Error).Methods("POST")
	//r.HandleFunc("", show404Error).Methods("")
	r.NotFoundHandler = http.HandlerFunc(show404Error)
	log.Fatal(http.ListenAndServe(":4000", r))
}
func show405Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func increaseRegCount(w http.ResponseWriter, r *http.Request) {

	xRegCount += 1
	xRegCountString := strconv.Itoa(xRegCount)
	w.Header().Set("x-Reg-Count", xRegCountString)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ping)

}

func show404Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
*/

package main

func main() {

}

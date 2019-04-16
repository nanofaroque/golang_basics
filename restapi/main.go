package restapi

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// our interfaces function
func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

	// Write the router
	router.Handle("/people",GetPeople).Methods("GET")
	router.Handle("/people/{id}",GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

}
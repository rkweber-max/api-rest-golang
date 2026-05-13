package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkweber-max/api-rest-golang/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.GitById).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8090", r))
}

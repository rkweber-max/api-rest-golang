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
	r.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8090", r))
}

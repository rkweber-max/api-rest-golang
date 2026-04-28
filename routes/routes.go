package routes

import (
	"log"
	"net/http"

	"github.com/rkweber-max/api-rest-golang/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

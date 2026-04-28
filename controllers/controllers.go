package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rkweber-max/api-rest-golang/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

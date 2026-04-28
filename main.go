package main

import (
	"fmt"

	"github.com/rkweber-max/api-rest-golang/models"
	"github.com/rkweber-max/api-rest-golang/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Test 1", History: "Teste 12"},
		{Name: "Test 2", History: "Teste 21"},
	}

	fmt.Println("Starting rest serve in Go")
	routes.HandleRequest()
}

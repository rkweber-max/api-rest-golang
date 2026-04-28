package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
	log.Fatal(":8000", nil)
}

func HandleRequest() {
	http.HandleFunc("/", Home)
}

func main() {
	fmt.Println("Starting rest serve in Go")
	HandleRequest()
}

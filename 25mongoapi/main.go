package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arc1997/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb API")
	fmt.Println("Server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8000", r))
}

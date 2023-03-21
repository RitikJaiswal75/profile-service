package main

import (
	"fmt"
	"log"
	"net/http"
	"profile-service/routes/index"
)

func main() {
	index.BaseHandler()
	fmt.Println("Starting server at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

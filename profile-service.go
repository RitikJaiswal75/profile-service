package main

import (
	"fmt"
	"log"
	"net/http"
	"profile-service/routes/index"
)

func main() {
	index.BaseHandler()
	fmt.Println("Starting server at 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

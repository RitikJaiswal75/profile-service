package health

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Supported.", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Profile Service - Health Verified")
}

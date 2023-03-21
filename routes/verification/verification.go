package verification

import (
	"encoding/json"
	"log"
	"net/http"

	"profile-service/utils"

	"golang.org/x/crypto/bcrypt"
)

type ReqBody struct {
	Salt string `json:"salt"`
}

type ResBody struct {
	Hash string `json:"hash"`
}

func VerificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/verification" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method Not Supported.", http.StatusNotFound)
		return
	}

	var req ReqBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Fatal(err)
	}

	if req.Salt == "" {
		http.Error(w, "Salt not found", http.StatusNotFound)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(utils.GoDotEnvVariable("CHAIN_CODE")), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Error while encryption", http.StatusInternalServerError)
		return
	}

	var res ResBody
	res.Hash = string(hash)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

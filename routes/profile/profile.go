package profile

import (
	"encoding/json"
	"net/http"
	"strconv"

	"profile-service/utils"
)

type Profile struct {
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Yoe          int    `json:"yoe"`
	Company      string `json:"company"`
	Designation  string `json:"designation"`
	Github_id    string `json:"github_id"`
	Linkedin_id  string `json:"linkedin_id"`
	Twitter_id   string `json:"twitter_id"`
	Instagram_id string `json:"instagram_id"`
	Website      string `json:"website"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported.", http.StatusNotFound)
	}

	var res Profile

	res.First_name = utils.GoDotEnvVariable("FIRST_NAME")
	res.Last_name = utils.GoDotEnvVariable("LAST_NAME")
	res.Email = utils.GoDotEnvVariable("EMAIL")
	res.Phone = utils.GoDotEnvVariable("PHONE")
	res.Yoe, _ = strconv.Atoi(utils.GoDotEnvVariable("YOE"))
	res.Company = utils.GoDotEnvVariable("COMPANY")
	res.Designation = utils.GoDotEnvVariable("DESIGNATION")
	res.Github_id = utils.GoDotEnvVariable("GITHUB_ID")
	res.Linkedin_id = utils.GoDotEnvVariable("LINKEDIN_ID")
	res.Twitter_id = utils.GoDotEnvVariable("TWITTER_ID")
	res.Instagram_id = utils.GoDotEnvVariable("INSTAGRAM_ID")
	res.Website = utils.GoDotEnvVariable("WEBSITE")

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		if err != nil {
			http.Error(w, "Error getting profile data", http.StatusInternalServerError)
		}
	}
}

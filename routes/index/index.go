package index

import (
	"net/http"
	"profile-service/routes/health"
	"profile-service/routes/profile"
	"profile-service/routes/verification"
)

func BaseHandler() {
	http.HandleFunc("/health", health.HealthHandler)

	http.HandleFunc("/verification", verification.VerificationHandler)

	http.HandleFunc("/profile", profile.ProfileHandler)
}

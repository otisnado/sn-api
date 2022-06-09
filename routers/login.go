package routers

import (
	"encoding/json"
	"net/http"

	"github.com/otisnado/sn-api/db"
	"github.com/otisnado/sn-api/jwt"
	"github.com/otisnado/sn-api/models"
)

/*Login user*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or password is invalid "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "User email is required ", 400)
	}
	document, exists := db.LoginAttempt(t.Email, t.Password)
	if !exists {
		http.Error(w, "User or password is invalid", 400)
	}
	jwtKey, err := jwt.JWTGenerator(document)
	if err != nil {
		http.Error(w, "JWT generation error "+err.Error(), 400)
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

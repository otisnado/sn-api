package routers

import (
	"encoding/json"
	"net/http"

	"github.com/otisnado/sn-api/db"
	"github.com/otisnado/sn-api/models"
)

/*Register is to create new users */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password len must be over 6 char", 400)
	}
	_, found, _ := db.CheckUserExists(t.Email)
	if found == true {
		http.Error(w, "An user already exist with given email", 400)
		return
	}

	_, status, err := db.InsertRegistry(t)
	if err != nil {
		http.Error(w, "Unsuccesfull insert of data "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Unsuccessfull insert of user data", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

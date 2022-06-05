package middlew

import (
	"net/http"

	"github.com/otisnado/sn-api/db"
)

/*CheckDB function for checking database connection state*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Database connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

package logins

import (
	"encoding/json"
	"net/http"
)

// GetBeers returns the cellar
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}

package logins

import (
	"encoding/json"
	//"io/ioutil"
	"net/http"
	"../../internal"
	"../../api/customAuth"
)

// GetBeers returns the cellar
//func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode("")
//}

func LoginUser(writer http.ResponseWriter, request *http.Request) {
	var userRequest customAuth.FirebaseUserRequest

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userRequest)

	if err != nil {
		http.Error(writer, "Login Failed", http.StatusNotAcceptable)
		return
	}
	internal.Auth()

	userInfo, response := customAuth.SigninWithPassword(userRequest)
	if response != nil {
		http.Error(writer, response.Status, response.StatusCode)
		return
	}

	_userInfo, err := json.Marshal(userInfo)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(_userInfo)
	//auth.
}

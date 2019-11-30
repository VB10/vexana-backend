package customAuth

import (
	"../../configs"
	"bytes"
	"encoding/json"
	"firebase.google.com/go/auth"
	"net/http"
)

func SigninWithPassword(userRequest FirebaseUserRequest) (*auth.UserRecord, *http.Response) {
	_userRequest, err := json.Marshal(userRequest)
	if err != nil {
		return nil, &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusNotAcceptable,
		}
	}
	//var apiKey = app

	resp, err := http.Post(configs.FIREBASE_AUTH_ENDPOINT, "application/json", bytes.NewBuffer(_userRequest))
	if err != nil {
		return nil, resp
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		var userFirebase auth.UserRecord
		body := json.NewDecoder(resp.Body).Decode(userFirebase)
		print(body)
		return &userFirebase, nil
	default:
		return nil, resp
	}

}

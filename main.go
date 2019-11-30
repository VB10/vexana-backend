package main

import (
	//"./internal"
	//"fmt"
	"github.com/gorilla/mux"
	"./pkg/logins"
	//"io"
	//"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", logins.LoginUser).Methods(http.MethodPost)
	_ = http.ListenAndServe(":8080", r)
}

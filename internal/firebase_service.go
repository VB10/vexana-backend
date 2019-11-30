package internal

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func internal() {
	print("init firebase")
	app, err := firebaseInit()
	if err != nil {
		return
	}
	firebaseApp = app
}

func firebaseInit() (*firebase.App, error) {
	opt := option.WithCredentialsFile("/Users/a02484320/Documents/firebase/vbtest_golang_conf.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}


func firebaseConfig() (*firebase.App, error) {
	opt := option.WithCredentialsFile("/Users/a02484320/Documents/firebase/vbtest_golang_conf.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}



func Auth() (*auth.Client, error) {
	if firebaseApp == nil {
		app, err := firebaseInit()
		if err != nil {
			return nil, err
		}
		firebaseApp = app
	}

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing auth: %v", err)
	}

	u, err := auth.GetUserByEmail(context.Background(), "veli123@hwa.com")

	if err != nil {
		print(err)
	}

	print(u)
	return auth, nil
}

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// GetCommitMessage
func (a *App) GetHello() string {
	// TODO HTTP call
	requestURL := fmt.Sprintf("http://localhost:%d/hello", 8081)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	} else {
		//fmt.Printf("client: got response!\n")
		//fmt.Printf("client: status code: %d\n", res.StatusCode)
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
		} else {
			return string(resBody)
		}
	}
	message := "NULL"
	return message
}

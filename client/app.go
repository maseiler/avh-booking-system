package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// Config struct
type AvhbsConfig struct {
	url        string
	clientName string
}

// User struct
type User struct { // TODO redundant (./client/models/User)
	Id        int
	Creation  time.Time
	Client    string
	BierName  string
	FirstName string
	LastName  string
	BoatName  string
	Status    string
	Email     string
	Phone     string
	Balance   float32
	MaxDebt   int
}

var avhbsConfig = AvhbsConfig{}

// NewApp creates a new App application struct
func NewApp(config AvhbsConfig) *App {
	avhbsConfig = config
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

// Get all users
func (a *App) GetAllUsers() []User {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := client.Get(avhbsConfig.url + "/getAllUsers")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	} else {
		//fmt.Printf("client: status code: %d\n", res.StatusCode)
		var users []User = UnmarshalUserList(res.Body)
		return users
	}
	return nil
}

// Returns the this client name
func (a *App) GetClientName() string {
	return avhbsConfig.clientName
}

// Adds new user
func (a *App) AddNewUser(newUser User) bool {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // TODO
			},
		},
	}

	newUser.Client = avhbsConfig.clientName
	userAsJson, err := json.Marshal(newUser)
	HandleMarshalingError(err)
	_, err = client.Post(avhbsConfig.url+"/addNewUser", "application/json", bytes.NewBuffer(userAsJson))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	} else {
		//fmt.Printf("client: status code: %d\n", res.StatusCode)
		//fmt.Printf("client: res: %s\n", res.Body)
		return true
	}
	return false
}

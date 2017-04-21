package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Config struct {
	Title               string `json:"title"`
	Consumer_key        string `json:"consumer_key"`
	Consumer_secret     string `json:"consumer_secret"`
	Access_token_key    string `json:"access_token_key"`
	Access_token_secret string `json:"access_token_secret"`
}
type Flock struct {
	ScreenNames []string
	Clients     []*twitter.Client
}

func readConfigTokensAndConnect() (flock Flock) {
	var config []Config
	var clients []*twitter.Client

	file, e := ioutil.ReadFile("flockConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
	fmt.Println("flockConfig.json read comlete")

	fmt.Print("connecting to twitter api --> ")
	for i := 0; i < len(config); i++ {
		configu := oauth1.NewConfig(config[i].Consumer_key, config[i].Consumer_secret)
		token := oauth1.NewToken(config[i].Access_token_key, config[i].Access_token_secret)
		httpClient := configu.Client(oauth1.NoContext, token)
		// twitter client
		client := twitter.NewClient(httpClient)
		clients = append(clients, client)
		flock.ScreenNames = append(flock.ScreenNames, config[i].Title)
	}
	flock.Clients = clients

	fmt.Println("connection successfull")

	return flock
}

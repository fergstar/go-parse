package main

import (
	"encoding/json"
	"fmt"
	"github.com/fergstar/go-parse/parse"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type config struct {
	ParseApplicationID string `envconfig:"PARSE_APPLICATION_ID"`
	ParseRestAPIKey    string `envconfig:"PARSE_REST_API_KEY"`
}

type GameScore struct {
	ObjectID string
	Name     string
}

func main() {
	var c config
	err := envconfig.Process("PARSE", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	if c.ParseApplicationID == "" || c.ParseRestAPIKey == "" {
		log.Fatal("Missing required environment variable")
	}

	http_client := &http.Client{}

	// Parse

	objectBody := &GameScore{Name: "Test"}

	client := parse.NewClient(http_client, c.ParseApplicationID, c.ParseRestAPIKey)

	success, _, err := client.Objects.Create("gamescore", objectBody)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(success.ObjectID)

	jsonRawMessage, _, err := client.Objects.Retrieve("gamescore", success.ObjectID)

	if err != nil {
		log.Fatal(err.Error())
	}

	var gameScore GameScore
	err = json.Unmarshal([]byte(*jsonRawMessage), &gameScore)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(gameScore)
}

package main

import (
	"net/http"

	"github.com/fergstar/go-parse/parse"
)

type config struct {
	ParseApplicationID string `envconfig:"APPLICATION_ID"`
	ParseRestAPIKey    string `envconfig:"RESTAPI_KEY"`
}

func main() {

	http_client := &http.Client{}

	// Parse

	client := parse.NewClient(http_client)

}

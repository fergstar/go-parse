package main

import (
	"net/http"

	"github.com/fergstar/go-parse/parse"
)

func main() {

	http_client := &http.Client{}

	// Parse

	client := parse.NewClient(http_client)

}

package parse

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ParseAPI required protocol and domain
const ParseAPI = "https://api.parse.com/1/"

// Client is a Parse client for making Parse API requests.
type Client struct {
	sling *sling.Sling

	// Different Parse API Services
	Objects *ObjectService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, ParseApplicationId, ParseRestApiKey string) *Client {

	base := sling.New().Client(httpClient).Base(ParseAPI).Set("X-Parse-Application-Id", ParseApplicationId).Set("X-Parse-REST-API-Key", ParseRestApiKey)

	return &Client{
		Objects: NewObjectService(base.New()),
	}
}

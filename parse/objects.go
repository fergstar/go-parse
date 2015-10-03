package parse

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// https://parse.com/docs/rest#objects

// ObjectService provides methods for accessing Parse object API endpoints.
type ObjectService struct {
	sling *sling.Sling
}

// NewObjectService returns a new ObjectService
func NewObjectService(sling *sling.Sling) *ObjectService {
	return &ObjectService{
		sling: sling.Path("classes/"),
	}
}

// Success represents a Parse API Object Succes response
type Success struct {
	CreatedAt string `json:"createdAt"`
	ObjectID  string `json:"objectId"`
}

// APIError represents a Parse API Error response
type APIError struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

// Create a new object on the Parse Cloud.
func (o *ObjectService) Create(className string, objectBody interface{}) (*Success, *http.Response, error) {

	success := new(Success)
	apiError := new(APIError)
	path := fmt.Sprintf("/%s", className)
	resp, err := o.sling.New().Post(path).BodyJSON(objectBody).Receive(success, apiError)
	return success, resp, err
}

/*
// Retrieve an object from the Parse Cloud.
func (o *ObjectService) Retrieve(className string) {

}

// Update the data on an object that already exists.
func (o *ObjectService) Update(className string) {

}

// Delete an object from the Parse Cloud.
func (o *ObjectService) Delete(className string) {

}
*/

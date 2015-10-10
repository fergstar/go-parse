package parse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

// SuccessCreate represents a Parse API Object Succes response
type SuccessCreate struct {
	CreatedAt string `json:"createdAt"`
	ObjectID  string `json:"objectId"`
}

// SuccessUpdate represents a Parse API Object Update successful update response
type SuccessUpdate struct {
	UpdatedAt time.Time `json:"updatedAt"`
}

// Create a new object on the Parse Cloud.
func (o *ObjectService) Create(className string, objectBody interface{}) (*SuccessCreate, *http.Response, error) {

	success := new(SuccessCreate)
	apiError := new(APIError)
	path := fmt.Sprintf("%s", className)
	resp, err := o.sling.New().Post(path).BodyJSON(objectBody).Receive(success, apiError)
	return success, resp, releventError(err, apiError)
}

// Retrieve an object from the Parse Cloud.
func (o *ObjectService) Retrieve(className, objectID string) (*json.RawMessage, *http.Response, error) {

	var parseObject *json.RawMessage = &json.RawMessage{}
	apiError := new(APIError)
	path := fmt.Sprintf("%s/%s", className, objectID)
	resp, err := o.sling.New().Get(path).Receive(parseObject, apiError)

	return parseObject, resp, releventError(err, apiError)
}

// Update the data on an object that already exists.
func (o *ObjectService) Update(className, objectID string, objectBody interface{}) (*SuccessUpdate, *http.Response, error) {
	success := new(SuccessUpdate)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/%s", className, objectID)
	resp, err := o.sling.New().Put(path).BodyJSON(objectBody).Receive(success, apiError)

	return success, resp, releventError(err, apiError)
}

// Delete an object from the Parse Cloud.
func (o *ObjectService) Delete(className, objectID string) (*http.Response, error) {

	apiError := new(APIError)
	path := fmt.Sprintf("%s/%s", className, objectID)
	resp, err := o.sling.New().Delete(path).Receive(nil, apiError)

	return resp, releventError(err, apiError)
}

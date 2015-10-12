package parse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// https://parse.com/docs/rest#users

// UserService provides methods for accessing Parse user API endpoints.
type UserService struct {
	sling *sling.Sling
}

// NewUserService return a new UserService
func NewUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path(""),
	}
}

// Success when a user succesfully signsup
type SuccessSignup struct {
	CreatedAt    string `json:"createdAt"`
	ObjectID     string `json:"objectId"`
	SessionToken string `json:"sessionToken"`
}

// SignUp a new user
func (u *UserService) SignUp(userDetails interface{}) (*SuccessSignup, *http.Response, error) {
	// TODO marshal userDetails into user struct and
	// validate username and password available

	success := new(SuccessSignup)
	apiError := new(APIError)
	path := fmt.Sprintf("users")
	resp, err := u.sling.New().Post(path).BodyJSON(userDetails).Receive(success, apiError)

	return success, resp, releventError(err, apiError)
}

// LogIn user
func (u *UserService) LogIn(userDetails interface{}) (*json.RawMessage, *http.Response, error) {

	var successLogin = &json.RawMessage{}
	apiError := new(APIError)
	path := fmt.Sprintf("login")
	resp, err := u.sling.New().Post(path).Receive(successLogin, apiError)

	return successLogin, resp, releventError(err, apiError)
}

// Delete user from Parse
func (u *UserService) Delete(objectID, sessionToken string) (*http.Response, error) {

	apiError := new(APIError)
	path := fmt.Sprintf("users/%s", objectID)
	resp, err := u.sling.New().Set("X-Parse-Session-Token", sessionToken).Delete(path).Receive(nil, apiError)

	return resp, releventError(err, apiError)
}

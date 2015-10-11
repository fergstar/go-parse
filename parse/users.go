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
type Success struct {
	CreatedAt    string `json:"createdAt"`
	ObjectID     string `json:"objectId"`
	SessionToken string `json:"sessionToken"`
}

// SignUp a new user
func (u *UserService) SignUp(userDetails interface{}) (*Success, *http.Response, error) {
	// TODO marshal userDetails into user struct and
	// validate username and password available

	success := new(Success)
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

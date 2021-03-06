package parse

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type GameScore struct {
	Score      int
	PlayerName string
	CheatMode  bool
}

func TestObjectsService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1/classes/{className}", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		fmt.Println(r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"createdAt":"2011-08-20T02:06:57.931Z","objectId":"Ed1nuqPvcm"}`)
	})

	testObject := &GameScore{Score: 23, PlayerName: "Test", CheatMode: false}

	client := NewClient(httpClient, "ApplicationId", "RestAPIkey")
	success, _, err := client.Objects.Create("gamescore", testObject)

	if err != nil {
		t.Errorf("Objects.Create error %v", err)
	}

	expected := &Success{CreatedAt: "2011-08-20T02:06:57.931Z", ObjectID: "Ed1nuqPvcm"}

	if !reflect.DeepEqual(expected, success) {
		t.Errorf("Objects.Create expected:\n%+v, got:\n %+v", expected, success)
	}

}

func TestObjectsService_Retrieve(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1/classes/{className}/{objectId}", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
	})

	client := NewClient(httpClient, "ApplicationId", "RestAPIkey")
	jsonMessage, _, err := client.Objects.Retrieve("gamescore", "Ed1nuqPvcm")

	if err != nil {
		t.Errorf("Objects.Retrieve error %v", err)
	}

	expected := &GameScore{Score: 23, PlayerName: "Test", CheatMode: false}

	if !reflect.DeepEqual(expected, jsonMessage) {
		t.Errorf("Objects.Retrieve expected:\n%+v, got:\n %+v", expected, jsonMessage)
	}

}

// func TestObjectService_Update

// func TestObjectService_Delete

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

	mux.HandleFunc("/1/classes/gamescore", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		fmt.Println(r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"createdAt":"2011-08-20T02:06:57.931Z","objectId":"Ed1nuqPvcm"}`)
	})

	testObject := &GameScore{Score: 23, PlayerName: "Test", CheatMode: false}

	client := NewClient(httpClient)
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

	client := NewClient(httpClient)
}

// func TestObjectService_Update

// func TestObjectService_Delete

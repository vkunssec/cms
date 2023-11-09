package main_test

import (
	"cms/pkg/handler"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(t *testing.T) {
	expected := handler.ResponseIndex{
		Message: "OK",
		Version: "1.0.0",
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["message"] != expected.Message {
		t.Errorf("message invalid")
	}
	if response["version"] != expected.Version {
		t.Errorf("version invalid")
	}
}

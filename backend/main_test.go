package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", index)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/index", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}

	var responseJson Response
	json.Unmarshal(writer.Body.Bytes(), &responseJson)
	if responseJson.Name != "TestName" {
		t.Errorf("Response name is %v", responseJson.Name)
	}
	if responseJson.Message != "TestMessage" {
		t.Errorf("Response message is %v", responseJson.Message)
	}
}

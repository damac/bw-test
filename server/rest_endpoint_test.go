package main

import (
	"brightwheel-interview/proto/device_rest"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Return 405 on GET to POST on /v1/reading
func TestBadPost(t *testing.T) {
	t.Run("It returns 400 when POST is used instead of GET.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/v1/reading", nil)
		response := httptest.NewRecorder()
		receiveReadings(response, request)
		if response.Code != http.StatusMethodNotAllowed {
			t.Error("Expected 405, I did not receive that.")
		}
	})
}

// 404 on requesting count for an absent ID
func TestAbsentId(t *testing.T) {
	t.Run("It returns 404 when a blank, ID is requested.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/v1/latest/", nil)
		response := httptest.NewRecorder()
		returnCount(response, request)
		if response.Code != http.StatusNotFound {
			t.Error("Expected a 404 response, I did not receive that.")
		}
	})
}

// 404 on a non-existent ID
func TestMissingId(t *testing.T) {
	t.Run("It returns 404 when a non-existent ID is requested.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/v1/latest/hello-other-engineers-how-are-you", nil)
		response := httptest.NewRecorder()
		returnCount(response, request)
		if response.Code != http.StatusNotFound {
			t.Error("Expected a 404 response, I did not receive that.")
		}
	})
}

// Validate protobuf is lining up with json examples
func TestUnmarshal(t *testing.T) {
	var testMessage = &device_rest.SomeMessage{}
	data, err := os.ReadFile("../examples/example.json")
	if err != nil {
		t.Error("Unable to read json examples file.")
	}
	errJson := json.Unmarshal(data, testMessage)
	if errJson != nil {
		t.Error("Unable to unmarshal json examples file.")
	}
	t.Run("Data unmarshal examples", func(t *testing.T) {
		if testMessage.Id != "36d5658a-6908-479e-887e-a949ec199272" {
			t.Error("Failed to unmarshal json, check examples/example.json and protobuf")
		}
	})
}

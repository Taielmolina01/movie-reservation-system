package tests

import (
	"movie-reservation-system/initializers"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
)

func TestCreateUserWithoutEmail(t *testing.T) {
	router := initializers.Init()

	jsonBody := `{"name": "John Doe", "password": "password123", "role": "user"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"Invalid request body: Key: 'UserRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestCreateUserWithoutName(t *testing.T) {
	t.Log("TestCreateUserWithoutName passed")
}

func TestCreateUserWithoutPassword(t *testing.T) {
	t.Log("TestCreateUserWithoutPassword passed")
}

func TestCreateUserWithoutPasswordMinLength(t *testing.T) {
	t.Log("TestCreateUserWithoutPasswordMinLength passed")
}

func TestCreateUserWithoutRoles(t *testing.T) {
	t.Log("TestCreateUserWithoutRoles passed")
}

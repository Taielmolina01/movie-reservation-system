package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserWithoutEmail(t *testing.T) {
	t.Log("Try to create an user without email")
	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Name": "John Doe", "Email": "", "Password": "password123", "Role": "user"}`

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

func TestCreateUserWithInvalidEmail(t *testing.T) {
	t.Log("Try to create an user with a not email in the email field")

	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Name": "John Doe", "Email": "invalid-email", "Password": "password123", "Role": "user"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"Invalid request body: Key: 'UserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestCreateUserWithoutName(t *testing.T) {
	t.Log("Try to create an user without name")

	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Name": "", "Email": "johndoe@gmail.com", "Password": "password123", "Role": "user"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"Invalid request body: Key: 'UserRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestCreateUserWithoutPassword(t *testing.T) {
	t.Log("Try to create an user without password")

	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Name": "John Doe", "Email": "johndoe@gmail.com", "Password": "", "Role": "user"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"Invalid request body: Key: 'UserRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestCreateUserWithoutPasswordMinLength(t *testing.T) {
	t.Log("Try to create an user with a password's length less than eight")

	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Name": "John Doe", "Email": "johndoe@gmail.com", "Password": "short", "Role": "user"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"Invalid request body: Key: 'UserRequest.Password' Error:Field validation for 'Password' failed on the 'min' tag"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestCreateUserWithoutRoles(t *testing.T) {
	t.Log("Try to create an user without role")

	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	/*
		expected := `{"message":"user"}`
		if recorder.Body.String() != expected {
			t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
		}
	*/
}

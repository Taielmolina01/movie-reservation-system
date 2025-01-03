package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginNonExistentUser(t *testing.T) {
	t.Log("Try to login with an user that is not registered")
	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Email": "johndoe@gmail.com", "Password": "myPassword"}`

	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestLoginExistentUserWithWrongPassword(t *testing.T) {
	t.Log("Try to login with an user with wrong password")

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

	secondJsonBody := `{"Email": "johndoe@gmail.com", "Password": "anotherPassword"}`

	secondReq, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(secondJsonBody))
	secondReq.Header.Set("Content-Type", "application/json")

	secondRecorder := httptest.NewRecorder()
	router.ServeHTTP(secondRecorder, secondReq)

	if secondRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, secondRecorder.Code)
	}

	expected := `{"error":"The entered password is not the user's password"}`
	if secondRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
	}
}

func TestLoginWithCorrectFields(t *testing.T) {
	t.Log("Try to login with a registered user with correct password")

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

	secondJsonBody := `{"email": "johndoe@gmail.com", "password": "myPassword"}`

	secondReq, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(secondJsonBody))
	secondReq.Header.Set("Content-Type", "application/json")

	secondRecorder := httptest.NewRecorder()
	router.ServeHTTP(secondRecorder, secondReq)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}

	/*
		expected := `{"message":"The entered password is not the user's password"}`
		if secondRecorder.Body.String() != expected {
			t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
		}
	*/

	// Im not so sure about how i should test the response token, so i only check the status for now
}

func TestLogoutNonExistentUser(t *testing.T) {
	t.Log("Try to logout with an non existent user")

	router, err := SetUpRouterTest()

	if err != nil {
		t.Error("Error creating test router")
	}

	jsonBody := `{"Email": "johndoe@gmail.com", "Password": "password123"}`

	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestLogoutWithANotLoggedUser(t *testing.T) {
	t.Log("Try to logout with a not logged user")

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

	secondReq, _ := http.NewRequest("POST", "/logout/johndoe@gmail.com", nil)
	secondReq.Header.Set("Content-Type", "application/json")

	secondRecorder := httptest.NewRecorder()
	router.ServeHTTP(secondRecorder, secondReq)

	if secondRecorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, secondRecorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com does not have a token"}`
	if secondRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
	}
}

func TestLogoutWithALoggedUser(t *testing.T) {
	t.Log("Try to logout with a logged user")

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

	secondJsonBody := `{"email": "johndoe@gmail.com", "password": "myPassword"}`

	secondReq, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(secondJsonBody))
	secondReq.Header.Set("Content-Type", "application/json")

	secondRecorder := httptest.NewRecorder()
	router.ServeHTTP(secondRecorder, secondReq)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}

	thirdReq, _ := http.NewRequest("POST", "/logout/johndoe@gmail.com", nil)
	thirdReq.Header.Set("Content-Type", "application/json")

	thirdRecorder := httptest.NewRecorder()
	router.ServeHTTP(thirdRecorder, thirdReq)

	if thirdRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, thirdRecorder.Code)
	}

	/*
		expected := `{"error":"User with email johndoe@gmail.com does not have a token"}`
		if secondRecorder.Body.String() != expected {
			t.Errorf("Expected body %s but got %s", expected, thirdRecorder.Body.String())
		}
	*/

	// Same issue as with succesful login
}

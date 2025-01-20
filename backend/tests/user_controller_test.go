package tests

import (
	"net/http"
	"testing"
	"encoding/json"
)

// Tests of user's creations

func TestCreateUserWithoutEmail(t *testing.T) {
	t.Log("Try to create an user without email")
	router := UseRouter(t)

	jsonBody := `{"Name": "John Doe", "Email": "", "Password": "password123", "Role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

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

	router := UseRouter(t)

	jsonBody := `{"Name": "John Doe", "Email": "invalid-email", "Password": "password123", "Role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

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

	router := UseRouter(t)

	jsonBody := `{"Name": "", "Email": "johndoe@gmail.com", "Password": "password123", "Role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

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

	router := UseRouter(t)

	jsonBody := `{"Name": "John Doe", "Email": "johndoe@gmail.com", "Password": "", "Role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

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

	router := UseRouter(t)

	jsonBody := `{"Name": "John Doe", "Email": "johndoe@gmail.com", "Password": "short", "Role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

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

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}
}

func TestCreateUserWithAdminRole(t *testing.T) {
	t.Log("Try to create an user with admin role")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "admin"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}
}

func TestCreateUserWithUserRole(t *testing.T) {
	t.Log("Try to create an user with user role")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}
}

func TestCreateUserWithNonExistentRole(t *testing.T) {
	t.Log("Try to create an user with not existent role")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "adminuser"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}
}

// Test of user updates

// Hacer un test sin loggearse y despues todo lo demás loggeado

func TestUpdateANotExistentUser(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "admin"}`

	recorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com", jsonBody)

	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"Missing authentication token"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestUpdateUserWithoutName(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	accessToken, err := GetAccessToken(secondRecorder)

	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	secondJsonBody := `{"name":""}`

	req, _ := http.NewRequest("PUT", "/users/johndoe@gmail.com", secondJsonBody)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	
	secondRecorder := PerformRequestWithRequest(t, router, req)

	if secondRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"User must have a name"}`
	if secondRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
	}
}

func TestUpdateUserWithoutRole(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"role":""}`

	secondRecorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com", secondJsonBody)

	if secondRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	expected := `{"error":"Invalid request body: Key: 'UserUpdateRequest.Role' Error:Field validation for 'Role' failed on the 'oneof' tag"}`
	if secondRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
	}

	thirdJsonBody := `{"role": "adminuser"}`

	thirdRecorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com", thirdJsonBody)

	if thirdRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, recorder.Code)
	}

	if thirdRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, thirdRecorder.Body.String())
	}
}

func TestUpdateUserWithValidNameAndRole(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"name":"John Doe II", "role": "admin"}`

	secondRecorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com", secondJsonBody)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, secondRecorder.Code)
	}	

	var response struct {
		Message struct {
			Name  string `json:"Name"`
			Role  string `json:"Role"`
			Email string `json:"Email"`
		} `json:"message"`
	}
	err := json.Unmarshal(secondRecorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error decoding response JSON: %v", err)
	}

	if response.Message.Name != "John Doe II" {
		t.Errorf("Expected name 'John Doe II', but got '%v'", response.Message.Name)
	}

	if response.Message.Role != "admin" {
		t.Errorf("Expected role 'admin', but got '%v'", response.Message.Role)
	}
}

// Test of user password updates

// Hacer un test sin loggearse y despues todo lo demás loggeado

func TestUpdateUserPasswordOfUnregisteredUser(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"OldPassword": "myPassword", "NewPassword": "myNewPassword"}`

	recorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com/password", jsonBody)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestUpdateUserPasswordWithInvalidOldPassword(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"OldPassword": "notMyPassword", "NewPassword": "myAttempOfNewPassword"}`

	secondRecorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com/password", secondJsonBody)

	if secondRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, secondRecorder.Code)
	}
}

func TestUpdateUserPasswordWithLengthLessThanEight(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"OldPassword": "myPassword", "NewPassword": "short"}`

	secondRecorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com/password", secondJsonBody)

	if secondRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, secondRecorder.Code)
	}
}

func TestUpdateUserPasswordWithValidPassword(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"OldPassword": "myPassword", "NewPassword": "myLongPassword"}`

	secondRecorder := PerformRequest(t, router, "PUT", "/users/johndoe@gmail.com/password", secondJsonBody)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}
}

// Test of user delete

// Hacer un test sin loggearse y despues todo lo demás loggeado

func TestDeleteANonExistentUser(t *testing.T) {
	router := UseRouter(t)

	recorder := PerformRequest(t, router, "DELETE", "/users/johndoe@gmail.com", "")

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestDeleteAnExistentUser(t *testing.T) {
	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword", "role": "user"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondRecorder := PerformRequest(t, router, "GET", "/users/johndoe@gmail.com", "")

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}

	thirdRecorder := PerformRequest(t, router, "DELETE", "/users/johndoe@gmail.com", "")

	if thirdRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, thirdRecorder.Code)
	}

	fourthRecorder := PerformRequest(t, router, "GET", "/users/johndoe@gmail.com", "")

	if fourthRecorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, fourthRecorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if fourthRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, fourthRecorder.Body.String())
	}
}
package balance

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetBalance_Success tests the GetBalance function with a successful HTTP response.
func TestGetBalance_Success(t *testing.T) {
	// Mock response
	mockResponse := `{"balance":"1000","balanceHint":"1000","lockedBalance":"500","lockedBalanceHint":"500","utxoNum":1}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	address := "testAddress"
	expectedBalanceHint := "1000"
	actualBalanceHint, err := GetBalance(address, func(addr string) string {
		return server.URL
	})

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if actualBalanceHint != expectedBalanceHint {
		t.Errorf("Expected balance hint %s but got %s", expectedBalanceHint, actualBalanceHint)
	}
}

// TestGetBalance_HttpError tests the GetBalance function when an HTTP error occurs.
func TestGetBalance_HttpError(t *testing.T) {
	// Mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	address := "testAddress"

	_, err := GetBalance(address, func(addr string) string {
		return server.URL
	})

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	expectedErrorMsg := "received non-200 HTTP status"
	if !containsErrorMessage(err, expectedErrorMsg) {
		t.Errorf("Expected error message to contain %q but got %v", expectedErrorMsg, err)
	}
}

// TestGetBalance_UnmarshalError tests the GetBalance function when JSON unmarshalling fails.
func TestGetBalance_UnmarshalError(t *testing.T) {
	// Mock response with invalid JSON
	mockResponse := `{"balance":1000,"balanceHint":1000` // Invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	address := "testAddress"
	_, err := GetBalance(address, func(addr string) string {
		return server.URL
	})

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	expectedErrorMsg := "failed to unmarshal JSON"
	if !containsErrorMessage(err, expectedErrorMsg) {
		t.Errorf("Expected error message to contain %q but got %v", expectedErrorMsg, err)
	}
}

// TestGetBalance_RequestError tests the GetBalance function when the HTTP request fails.
func TestGetBalance_RequestError(t *testing.T) {
	// Simulate a request error by providing an invalid URL
	_, err := GetBalance("testAddress", func(addr string) string {
		return "http://invalid-url"
	})

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	expectedErrorMsg := "failed to make HTTP request"
	if !containsErrorMessage(err, expectedErrorMsg) {
		t.Errorf("Expected error message to contain %q but got %v", expectedErrorMsg, err)
	}
}

// Utility function to check if the error message contains the expected substring
func containsErrorMessage(err error, expectedMessage string) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), expectedMessage)
}

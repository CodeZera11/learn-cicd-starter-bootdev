package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Error generating request: %s", err.Error())
	}

	req.Header.Set("Authorization", "ApiKey my-test-key")

	key, err := GetAPIKey(req.Header)

	if err != nil {
		t.Fatalf("Error getting api key: %s", err.Error())
	}

	if key == "" {
		t.Fatalf("Key not found!")
	}
}

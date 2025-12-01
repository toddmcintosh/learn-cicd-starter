package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_NoHeader(t *testing.T) {
	h := http.Header{}
	_, err := GetAPIKey(h)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_Malformed(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "Bearer xyz")

	_, err := GetAPIKey(h)

	if err == nil {
		t.Fatalf("expected malformed header error, got nil")
	}
}

func TestGetAPIKey_Success(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if key != "abc123" {
		t.Fatalf("expected abc123, got %s", key)
	}
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	webAscii "webAscii/server"
)

func TestAsciiServer_MethodNotAllowed(t *testing.T) {
	// Prepare a GET request
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	webAscii.AsciiServer(rec, req)

	// Check the response
	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status Method Not Allowed; got %v", resp.Status)
	}
}

func TestAsciiServer_BannerNotFound(t *testing.T) {
	// Prepare a POST request with an invalid banner
	body := strings.NewReader("Text=Hello&Banner=invalid")
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	webAscii.AsciiServer(rec, req)

	// Check the response
	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Not Found; got %v", resp.Status)
	}
}

func TestAsciiServer_ErrorParsingForm(t *testing.T) {
	// Prepare a POST request with malformed form data
	body := strings.NewReader("Text=Hello&Banner=")
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	webAscii.AsciiServer(rec, req)

	// Check the response
	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status Bad Request; got %v", resp.Status)
	}
}

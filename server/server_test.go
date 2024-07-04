package webAscii

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAsciiServer(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid POST request with input and banner",
			method:         http.MethodPost,
			url:            "/ascii",
			body:           "Input=Hello&Banner=default",
			expectedStatus: http.StatusOK,
			expectedBody:   "ASCII art output for 'Hello' with 'default' banner",
		},
		{
			name:           "POST request with missing input",
			method:         http.MethodPost,
			url:            "/ascii",
			body:           "Banner=default",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "ParseForm() error message",
		},
		{
			name:           "POST request with invalid banner",
			method:         http.MethodPost,
			url:            "/ascii",
			body:           "Input=Hello&Banner=invalid",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid banner specified",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				AsciiServer(w, r)
			})

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if body := strings.TrimSpace(rr.Body.String()); body != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got '%v' want '%v'",
					body, tt.expectedBody)
			}
		})
	}
}

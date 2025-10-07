package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid ApiKey header",
			headers: http.Header{
				"Authorization": []string{"ApiKey mykey123"},
			},
			expectedKey: "mykey123",
			expectError: false,
		},
		{
			name:        "missing Authorization header",
			headers:     http.Header{},
			expectError: true,
			errorMsg:    "no authorization header included",
		},
		{
			name: "malformed authorization header - only ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectError: true,
			errorMsg:    "malformed authorization header",
		},
		{
			name: "malformed authorization header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer token123"},
			},
			expectError: true,
			errorMsg:    "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error, but got none")
				} else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("expected error message '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
				}
				if key != tt.expectedKey {
					t.Errorf("expected key '%s', got '%s'", tt.expectedKey, key)
				}
			}
		})
	}
}

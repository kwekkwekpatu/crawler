package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name           string
		inputURL       string
		expected       string
		expectingError bool
	}{
		{
			name:           "remove scheme",
			inputURL:       "https://blog.boot.dev/path",
			expected:       "blog.boot.dev/path",
			expectingError: false,
		},
		{
			name:           "clear slash",
			inputURL:       "http://google.com/apples/",
			expected:       "google.com/apples",
			expectingError: false,
		},
		{
			name:           "already normalized",
			inputURL:       "blog.boot.dev/path",
			expected:       "blog.boot.dev/path",
			expectingError: false,
		},
		{
			name:           "invalid URL",
			inputURL:       "http://///invalid-url",
			expected:       "",
			expectingError: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if tc.expectingError {
				if err == nil {
					t.Errorf("Test %v - '%s' FAIL: expected error but got none", i, tc.name)
				}
				return
			}
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

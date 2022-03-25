package utils

import (
	"testing"
)

var test = []struct {
	key      string
	fallback string
}{
	{"PORT", "3000"},
	{"URL", "http://localhost:3000"},
}

func TestGetenv(t *testing.T) {
	for _, tt := range test {
		if value := Getenv(tt.key, tt.fallback); value != tt.fallback {
			t.Errorf("Getenv(%v, %v) = %v, want %v", tt.key, tt.fallback, value, tt.fallback)
		}
	}
}

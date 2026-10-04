package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetGreetingWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Ada", nil)
	rec := httptest.NewRecorder()

	GetGreeting(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var g Greeting
	if err := json.NewDecoder(rec.Body).Decode(&g); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if g.Name != "Ada" {
		t.Errorf("name = %q, want %q", g.Name, "Ada")
	}
	if !strings.Contains(g.Message, "Ada") {
		t.Errorf("message = %q, want it to address %q", g.Message, "Ada")
	}
}

func TestGetGreetingWithoutName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	GetGreeting(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var g Greeting
	if err := json.NewDecoder(rec.Body).Decode(&g); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if g.Message == "" {
		t.Error("message is empty, want a generic greeting")
	}
	if g.Name != "" {
		t.Errorf("name = %q, want empty", g.Name)
	}
}

func TestGetHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	GetHealth(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var h Health
	if err := json.NewDecoder(rec.Body).Decode(&h); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if h.Status == "" {
		t.Error("status is empty")
	}
}

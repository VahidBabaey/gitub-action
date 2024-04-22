package microservice

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler(res, req)
	expected := "Hello World!\n"
	if res.Body.String() != expected {
		t.Errorf("Unexpected body, got %v want %v", res.Body.String(), expected)
	}
}

package main

import (
	route "instagram_api/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(route.GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "data")
	}
}

func TestGetUser(t *testing.T) {
	// data :={"id":1,"name":"shivansh","email":"shivansh.singh20021@gmail.com","password":"ShivanshS"}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(route.GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "data")
	}
}

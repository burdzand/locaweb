package main

import (
	"github.com/andersondborges/locaweb/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMostRelevantsUrl(t *testing.T) {
	handler := http.HandlerFunc(controllers.ControllerMostRelevants)
	w := performRequest(handler, "GET", "/most_relevants/")

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if content := w.Header().Get("Content-Type"); content != "text/html; charset=utf-8" {
		t.Errorf("Content Type  returned wrong status code: got %v want text/html; charset=utf-8", content)
	}
}

func TestMostRelevantsJsonUrl(t *testing.T) {
	handler := http.HandlerFunc(controllers.ControllerMostRelevants)
	w := performRequest(handler, "GET", "/most_relevants/?export=json")

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if content := w.Header().Get("Content-Type"); content != "application/json" {
		t.Errorf("Content Type  returned wrong status code: got %v want application/json", content)
	}
}

func TestMostMentionedUsersUrl(t *testing.T) {
	handler := http.HandlerFunc(controllers.ControllerMostRelevants)
	w := performRequest(handler, "GET", "/most_mentions/")

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if content := w.Header().Get("Content-Type"); content != "text/html; charset=utf-8" {
		t.Errorf("Content Type  returned wrong status code: got %v want text/html; charset=utf-8", content)
	}
}

func TestMostMentionedUsersJsonUrl(t *testing.T) {
	handler := http.HandlerFunc(controllers.ControllerMostRelevants)
	w := performRequest(handler, "GET", "/most_mentions/?export=json")

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if content := w.Header().Get("Content-Type"); content != "application/json" {
		t.Errorf("Content Type  returned wrong status code: got %v want application/json", content)
	}
}

func TestLocawebApiUrl(t *testing.T) {

	req, err := http.NewRequest("GET", "http://tweeps.locaweb.com.br/tweeps", nil)

	client := &http.Client{}
	req.Header.Add("Username", "andersondborges@gmail.com")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	defer resp.Body.Close()
}

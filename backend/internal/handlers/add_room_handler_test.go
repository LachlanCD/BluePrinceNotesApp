package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddRoom(t *testing.T) {
	initTestingDB()

	expectedReturn := "4\n"
	expectedStatus := http.StatusCreated

	mux := http.NewServeMux()
	mux.HandleFunc("/api/room/{workspaceID}/add", AddRoom)

	form := url.Values{}
	form.Add("name", "Test")
	form.Add("colour", "Blue")

	req := httptest.NewRequest(http.MethodPost, "/api/room/{workspaceID}/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestAddRoomMissingColour(t *testing.T) {
	initTestingDB()

	expectedReturn := "Name and Colour must be populated\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/api/rooms/{workspaceID}/add", AddRoom)

	form := url.Values{}
	form.Add("name", "Test")

	req := httptest.NewRequest(http.MethodPost, "/api/rooms/test/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestAddRoomMissingName(t *testing.T) {
	initTestingDB()

	expectedReturn := "Name and Colour must be populated\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/api/room/{workspaceID}/add", AddRoom)

	form := url.Values{}
	form.Add("colour", "Blue")

	req := httptest.NewRequest(http.MethodPost, "/api/room/test/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestAddRoomMissingForm(t *testing.T) {
	initTestingDB()

	expectedReturn := "Name and Colour must be populated\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/api/room/{workspaceID}/add", AddRoom)

	req := httptest.NewRequest(http.MethodPost, "/api/room/test/add", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()

	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

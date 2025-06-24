package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestEditRoom(t *testing.T) {
	initTestingDB()

	expectedReturn := "{\"Id\":1,\"Name\":\"Test\",\"Colour\":\"Blue\",\"Notes\":\"\"}\n"
	expectedStatus := http.StatusOK

	mux := http.NewServeMux()
	mux.HandleFunc("/room/{id}/edit", EditRoom)

	form := url.Values{}
	form.Add("name", "Test")
	form.Add("colour", "Blue")

	req := httptest.NewRequest(http.MethodPost, "/room/1/edit", strings.NewReader(form.Encode()))
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

func TestEditRoomInvalidId(t *testing.T) {
	initTestingDB()

	expectedReturn := "Id must be a number\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/room/{id}/edit", EditRoom)

	form := url.Values{}
	form.Add("name", "Test")
	form.Add("colour", "Blue")

	req := httptest.NewRequest(http.MethodPost, "/room/t/edit", strings.NewReader(form.Encode()))
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

func TestEditRoomNotExist(t *testing.T) {
	initTestingDB()

	expectedReturn := "Unable to edit room\n"
	expectedStatus := http.StatusInternalServerError

	mux := http.NewServeMux()
	mux.HandleFunc("/room/{id}/edit", EditRoom)

	form := url.Values{}
	form.Add("name", "Test")
	form.Add("colour", "Blue")

	req := httptest.NewRequest(http.MethodPost, "/room/5/edit", strings.NewReader(form.Encode()))
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

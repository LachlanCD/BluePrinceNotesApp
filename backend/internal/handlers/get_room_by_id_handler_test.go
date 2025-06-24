package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRoomByIdBadRequest(t *testing.T) {
	initTestingDB()

	expectedReturn := "Id must be a number\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/rooms/{id}", GetRoomById)

	req := httptest.NewRequest(http.MethodGet, "/rooms/t", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestGetRoomByIdNotExist(t *testing.T) {
	initTestingDB()

	expectedReturn := "Unable to retrieve room\n"
	expectedStatus := http.StatusInternalServerError

	mux := http.NewServeMux()
	mux.HandleFunc("/rooms/{id}", GetRoomById)

	req := httptest.NewRequest(http.MethodGet, "/rooms/123", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestGetRoomById(t *testing.T) {
	initTestingDB()

	expectedReturn := "{\"Id\":1,\"Name\":\"room1\",\"Colour\":\"col1\",\"Notes\":\"note1\"}\n"
	expectedStatus := http.StatusOK

	mux := http.NewServeMux()
	mux.HandleFunc("/rooms/{id}", GetRoomById)

	url := "/rooms/1"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()

	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

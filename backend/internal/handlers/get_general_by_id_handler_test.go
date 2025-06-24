package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGeneralByIdBadRequest(t *testing.T) {
	initTestingDB()

	expectedReturn := "Id must be a number\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}", GetGeneralNoteById)

	req := httptest.NewRequest(http.MethodGet, "/general/j", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestGetGeneralByIdNotExist(t *testing.T) {
	initTestingDB()

	expectedReturn := "Unable to retrieve general note\n"
	expectedStatus := http.StatusInternalServerError

	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}", GetGeneralNoteById)

	req := httptest.NewRequest(http.MethodGet, "/general/123", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestGetGeneralById(t *testing.T) {
	initTestingDB()

	expectedReturn := "{\"Id\":1,\"Name\":\"gen1\",\"Notes\":\"note1\"}\n"
	expectedStatus := http.StatusOK

	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}", GetGeneralNoteById)

	url := "/general/1"
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

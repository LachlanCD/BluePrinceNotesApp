package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestEditGeneral(t *testing.T) {
	initTestingDB()

	expectedReturn := "{\"Id\":1,\"Name\":\"Test\",\"Notes\":\"\"}\n"
	expectedStatus := http.StatusOK

	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}/edit", EditGeneral)

	form := url.Values{}
	form.Add("name", "Test")

	req := httptest.NewRequest(http.MethodPost, "/general/1/edit", strings.NewReader(form.Encode()))
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

func TestEditGeneralInvalidId(t *testing.T) {
	initTestingDB()

	expectedReturn := "Id must be a number\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}/edit", EditGeneral)

	form := url.Values{}
	form.Add("name", "Test")

	req := httptest.NewRequest(http.MethodPost, "/general/t/edit", strings.NewReader(form.Encode()))
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

func TestEditGeneralNotExist(t *testing.T) {
	initTestingDB()

	expectedReturn := "Unable to edit general note\n"
	expectedStatus := http.StatusInternalServerError

	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}/edit", EditGeneral)

	form := url.Values{}
	form.Add("name", "Test")

	req := httptest.NewRequest(http.MethodPost, "/general/5/edit", strings.NewReader(form.Encode()))
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


package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddGeneral(t *testing.T) {
	initTestingDB()

	expectedStatus := http.StatusCreated
	expectedReturn := "3\n"

	mux := http.NewServeMux()
	mux.HandleFunc("/general/add", AddGeneralNote)

	form := url.Values{}
	form.Add("name", "Test")

	req := httptest.NewRequest(http.MethodPost, "/general/add", strings.NewReader(form.Encode()))
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

func TestAddGeneralMissingName(t *testing.T) {
	initTestingDB()

	expectedReturn := "Name must be populated\n"
	expectedStatus := http.StatusBadRequest
	
	mux := http.NewServeMux()
	mux.HandleFunc("/general/add", AddGeneralNote)

	form := url.Values{}

	req := httptest.NewRequest(http.MethodPost, "/general/add", strings.NewReader(form.Encode()))
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

func TestAddGeneralMissingForm(t *testing.T) {
	initTestingDB()

	expectedReturn := "Name must be populated\n"
	expectedStatus := http.StatusBadRequest
	
	mux := http.NewServeMux()
	mux.HandleFunc("/general/add", AddGeneralNote)
	mux.HandleFunc("/general/{id}/remove", RemoveGeneralById)

	req := httptest.NewRequest(http.MethodPost, "/general/add", nil)
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

func TestAddGeneralAlreadyExist(t *testing.T) {
	initTestingDB()

	expectedReturn := "Unable to add general note\n"
	expectedStatus := http.StatusInternalServerError
	
	mux := http.NewServeMux()
	mux.HandleFunc("/general/add", AddGeneralNote)
	mux.HandleFunc("/general/{id}/remove", RemoveGeneralById)

	form := url.Values{}
	form.Add("name", "gen1")

	req := httptest.NewRequest(http.MethodPost, "/general/add", strings.NewReader(form.Encode()))
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

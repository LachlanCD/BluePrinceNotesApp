package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoveGeneral(t *testing.T) {
	initTestingDB()

	expectedReturn := ""
	expectedStatus := http.StatusNoContent

	mux := http.NewServeMux()
	mux.HandleFunc("/api/general/{workspaceID}/{id}/remove", RemoveGeneralById)

	removeURL := "/api/general/" + "test/" + "1" + "/remove"
	req := httptest.NewRequest(http.MethodGet, removeURL, nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

func TestRemoveGeneralInvalidId(t *testing.T) {
	initTestingDB()

	expectedReturn := "Id must be a number\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/api/general/{workspaceID}/{id}/remove", RemoveRoomById)

	removeURL := "/api/general/" + "test/" + "t" + "/remove"
	req := httptest.NewRequest(http.MethodGet, removeURL, nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

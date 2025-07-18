package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoveRoom(t *testing.T) {
	initTestingDB()

	expectedReturn := ""
	expectedStatus := http.StatusNoContent

	mux := http.NewServeMux()
	mux.HandleFunc("/room/{id}/remove", RemoveRoomById)

	removeURL := "/room/" + "1" + "/remove"
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

func TestRemoveRoomInvalidId(t *testing.T) {
	initTestingDB()

	expectedReturn := "Id must be a number\n"
	expectedStatus := http.StatusBadRequest

	mux := http.NewServeMux()
	mux.HandleFunc("/room/{id}/remove", RemoveRoomById)

	removeURL := "/room/" + "t" + "/remove"
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

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllRooms(t *testing.T) {
	initTestingDB()
	expectedReturn := "[{\"Id\":1,\"Name\":\"room1\",\"Colour\":\"col1\",\"Notes\":\"\"},{\"Id\":2,\"Name\":\"room2\",\"Colour\":\"col2\",\"Notes\":\"\"},{\"Id\":3,\"Name\":\"room3\",\"Colour\":\"col3\",\"Notes\":\"\"}]\n"
	expectedStatus := http.StatusOK

	req := httptest.NewRequest(http.MethodGet, "/rooms", nil)
	w := httptest.NewRecorder()

	GetAllRooms(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

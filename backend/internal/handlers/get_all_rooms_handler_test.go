package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllRooms(t *testing.T) {
	expectedReturn := "[{\"Id\":1,\"Name\":\"room1\",\"Colour\":\"col1\",\"Notes\":\"\"},{\"Id\":2,\"Name\":\"room2\",\"Colour\":\"col2\",\"Notes\":\"\"},{\"Id\":3,\"Name\":\"room3\",\"Colour\":\"col3\",\"Notes\":\"\"}]\n"
	expectedStatus := http.StatusOK

	initTestingDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/rooms/{workspaceID}", GetAllRooms)

	req := httptest.NewRequest(http.MethodGet, "/api/rooms/test", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

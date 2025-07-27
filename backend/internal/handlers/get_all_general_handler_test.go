package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllGeneral(t *testing.T) {
	expectedStatus := http.StatusOK
	expectedReturn := "[{\"Id\":1,\"Name\":\"gen1\",\"Notes\":\"\"},{\"Id\":2,\"Name\":\"gen2\",\"Notes\":\"\"}]\n"

	initTestingDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/general/{workspaceID}", GetAllGeneral)

	req := httptest.NewRequest(http.MethodGet, "/api/general/test", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

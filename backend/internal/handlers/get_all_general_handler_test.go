package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllGeneral(t *testing.T) {
	initTestingDB()
	expectedStatus := http.StatusOK
	expectedReturn := "[{\"Id\":1,\"Name\":\"gen1\",\"Notes\":\"\"},{\"Id\":2,\"Name\":\"gen2\",\"Notes\":\"\"}]\n"

	req := httptest.NewRequest(http.MethodGet, "/general", nil)
	w := httptest.NewRecorder()

	GetAllGeneral(w, req)

	res := w.Result()
	checkStatus(expectedStatus, res.StatusCode, t)

	body := getBody(res, t)
	actualReturn := string(body)
	checkBody(expectedReturn, actualReturn, t)

	cleanDB()
}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllGeneral(t *testing.T) {
	initTestingDB()
	req := httptest.NewRequest(http.MethodGet, "/general", nil)
	w := httptest.NewRecorder()

	GetAllGeneral(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.StatusCode)
	}
}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllRooms(t *testing.T) {
	initTestingDB()
	req := httptest.NewRequest(http.MethodGet, "/rooms", nil)
	w := httptest.NewRecorder()

	GetAllRooms(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.StatusCode)
	}
}

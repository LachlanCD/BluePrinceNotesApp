package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetRoomByIdBadRequest(t *testing.T) {
	initTestingDB()
	req := httptest.NewRequest(http.MethodGet, "/rooms/0", nil)
	w := httptest.NewRecorder()

	GetRoomById(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400 Status Bad Request got %d", res.StatusCode)
	}
}

func TestGetRoomByIdNotExist(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/rooms/{id}", GetRoomById)

	req := httptest.NewRequest(http.MethodGet, "/rooms/123", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500 Status Internal Server Error got %d", res.StatusCode)
	}
}

func TestGetRoomById(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/rooms/add", AddRoom)
	mux.HandleFunc("/rooms/{id}", GetRoomById)
	mux.HandleFunc("/rooms/{id}/remove", RemoveRoomById)

	//--- Add Room ---
	form := url.Values{}
	form.Add("name", "Conference Room")
	form.Add("colour", "Blue")

	req1 := httptest.NewRequest(http.MethodPost, "/rooms/add", strings.NewReader(form.Encode()))
	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	res1 := w1.Result()
	defer res1.Body.Close()

	body, _ := io.ReadAll(res1.Body)
	roomID := strings.TrimSpace(string(body))
	t.Logf("Created room ID: %s", roomID)


	// --- Step 2: Get Room ---
	url := "/rooms/" + roomID
	req2 := httptest.NewRequest(http.MethodGet, url, nil)
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	res2 := w2.Result()

	if res2.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 Status Ok got %d", res2.StatusCode)
	}
	defer res2.Body.Close()
	body2, _ := io.ReadAll(res2.Body)
	t.Logf("Get room response: %s", body2)

	// --- Step 3: Remove Room ---
	removeURL := "/rooms/" + roomID + "/remove"
	req3 := httptest.NewRequest(http.MethodGet, removeURL, nil)
	w3 := httptest.NewRecorder()
	mux.ServeHTTP(w3, req3)

	res3 := w3.Result()
	if res3.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK on delete, got %d", res3.StatusCode)
	}
	defer res3.Body.Close()
	body3, _ := io.ReadAll(res3.Body)
	t.Logf("Remove room response: %s", body3)
}

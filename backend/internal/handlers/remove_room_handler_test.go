package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRemoveRoom(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/room/add", AddRoom)
	mux.HandleFunc("/room/{id}/remove", RemoveRoomById)

	//--- Add Note ---
	form := url.Values{}
	form.Add("name", "Test")
	form.Add("colour", "Blue")

	req1 := httptest.NewRequest(http.MethodPost, "/room/add", strings.NewReader(form.Encode()))
	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	res1 := w1.Result()
	defer res1.Body.Close()

	body, _ := io.ReadAll(res1.Body)
	noteID := strings.TrimSpace(string(body))

	if res1.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201 Status Created got %d", res1.StatusCode)
	}

	// --- Step 3: Remove Note ---
	removeURL := "/room/" + noteID + "/remove"
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

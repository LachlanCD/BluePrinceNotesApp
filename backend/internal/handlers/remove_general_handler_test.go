package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRemoveGeneral(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/general/add", AddGeneralNote)
	mux.HandleFunc("/general/{id}/remove", RemoveGeneralById)

	//--- Add Note ---
	form := url.Values{}
	form.Add("name", "Test")

	req1 := httptest.NewRequest(http.MethodPost, "/general/add", strings.NewReader(form.Encode()))
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
	removeURL := "/general/" + noteID + "/remove"
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

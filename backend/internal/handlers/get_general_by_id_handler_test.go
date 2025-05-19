package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetGeneralByIdBadRequest(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}", GetGeneralNoteById)

	req := httptest.NewRequest(http.MethodGet, "/general/j", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400 Status Bad Request got %d", res.StatusCode)
	}
}

func TestGetGeneralByIdNotExist(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/general/{id}", GetGeneralNoteById)

	req := httptest.NewRequest(http.MethodGet, "/general/123", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500 Status Internal Server Error got %d", res.StatusCode)
	}
}

func TestGetGeneralById(t *testing.T) {
	initTestingDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/general/add", AddGeneralNote)
	mux.HandleFunc("/general/{id}", GetGeneralNoteById)
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

	// --- Step 2: Get Note ---
	url := "/general/" +noteID 
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

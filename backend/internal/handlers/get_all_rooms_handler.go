package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	printRequest(r)

	data, err := db_interactions.ReadAllRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

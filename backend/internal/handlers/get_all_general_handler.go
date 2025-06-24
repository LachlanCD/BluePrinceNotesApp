package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func GetAllGeneral(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /general request\n")

	data, err := db_interactions.ReadAllGeneral()
	if err != nil {
		http.Error(w, "Unable to retrieve general notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

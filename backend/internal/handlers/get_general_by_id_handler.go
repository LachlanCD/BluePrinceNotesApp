package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func GetGeneralNoteById(w http.ResponseWriter, r *http.Request) {
	urlId:= r.PathValue("id")
	fmt.Printf("got /getroombyid/%s request\n", urlId)

	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusInternalServerError)
		return
	}

	data, err := db_interactions.ReadGeneralById(id)
	if err != nil {
		http.Error(w, "Unable to retrieve general note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}




package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
)

func AddGeneralNote(w http.ResponseWriter, r *http.Request) {
	printRequest(r)

	urlWorkspace := r.PathValue("workspaceID")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unabbble to parse form", http.StatusBadRequest)
		return
	}

	generalNote := models.General{
		Name: r.FormValue("name"),
	}

	if generalNote.Name == "" {
		http.Error(w, "Name must be populated", http.StatusBadRequest)
		return
	}

	data, err := db_interactions.AddGeneral(urlWorkspace, generalNote)
	if err != nil {
		http.Error(w, "Unable to add general note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

package handlers

import (
	"net/http"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func CreateWorkspaceID(w http.ResponseWriter, r *http.Request) {
	id, err := db_interactions.GenerateWorkspaceID()
	if err != nil {
		http.Error(w, "Error generating workspace ID", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
}

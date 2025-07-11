package handlers

import (
	"net/http"
	"strconv"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func RemoveRoomById(w http.ResponseWriter, r *http.Request) {
	printRequest(r)

	urlWorkspace := r.PathValue("workspaceID")

	urlId := r.PathValue("id")

	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusBadRequest)
		return
	}

	err = db_interactions.RemoveRoomNote(urlWorkspace, id)
	if err != nil {
		http.Error(w, "Unable to remove room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
)

func EditGeneral(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("id")
	fmt.Printf("got /getroombyid/%s request\n", urlId)

	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unabbble to parse form", http.StatusInternalServerError)
		return
	}

	generalNote := models.General{
		Id:     id,
		Name:   r.FormValue("name"),
		Notes:  r.FormValue("notes"),
	}

	err = db_interactions.UpdateGeneral(generalNote)
	if err != nil {
		http.Error(w, "Unable to add room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(generalNote)
}

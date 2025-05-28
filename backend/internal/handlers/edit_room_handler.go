package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
)

func EditRoom(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("id")
	fmt.Printf("got /getroombyid/%s request\n", urlId)

	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unabbble to parse form", http.StatusInternalServerError)
		return
	}

	room := models.Room{
		Id:     id,
		Name:   r.FormValue("name"),
		Colour: r.FormValue("colour"),
		Notes:  r.FormValue("notes"),
	}

	err = db_interactions.UpdateRoom(room)
	if err != nil {
		http.Error(w, "Unable to edit room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}


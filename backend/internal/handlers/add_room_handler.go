package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
)

func AddRoom(w http.ResponseWriter, r *http.Request) {
	printRequest(r)

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unabbble to parse form", http.StatusBadRequest)
		return
	}

	room := models.Room{
		Name:   r.FormValue("name"),
		Colour: r.FormValue("colour"),
	}

	if room.Name == "" || room.Colour == "" {
		http.Error(w, "Name and Colour must be populated", http.StatusBadRequest)
		return
	}

	data, err := db_interactions.AddRoom(room)
	if err != nil {
		http.Error(w, "Unable to add room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

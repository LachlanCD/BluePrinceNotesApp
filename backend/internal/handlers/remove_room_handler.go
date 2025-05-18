package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func RemoveRoomById(w http.ResponseWriter, r *http.Request) {
	urlId:= r.PathValue("id")
	fmt.Printf("got /getroombyid/%s request\n", urlId)

	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusInternalServerError)
		return
	}

	err = db_interactions.RemoveRoomNote(id)
	if err != nil {
		http.Error(w, "Unable to remove room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}



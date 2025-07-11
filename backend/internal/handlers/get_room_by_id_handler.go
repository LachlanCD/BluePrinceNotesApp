package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
)

func GetRoomById(w http.ResponseWriter, r *http.Request) {
	printRequest(r)

	urlWorkspace := r.PathValue("workspaceID")

	urlId := r.PathValue("id")

	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusBadRequest)
		return
	}

	data, err := db_interactions.ReadRoomById(urlWorkspace, id)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Unable to retrieve room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

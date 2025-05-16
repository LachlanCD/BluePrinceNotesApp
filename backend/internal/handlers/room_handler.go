package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}


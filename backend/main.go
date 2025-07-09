package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/handlers"

	"github.com/rs/cors"
)

func main() {
	// initialise db
	dbPath := "./data/notes.db"
	fmt.Println("running")
	err := db_interactions.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db created")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/rooms", handlers.GetAllRooms)
	mux.HandleFunc("/rooms/{id}", handlers.GetRoomById)
	mux.HandleFunc("/rooms/add", handlers.AddRoom)
	mux.HandleFunc("/rooms/{id}/update", handlers.EditRoom)
	mux.HandleFunc("/rooms/{id}/update/note", handlers.EditRoomNote)
	mux.HandleFunc("/rooms/{id}/remove", handlers.RemoveRoomById)
	mux.HandleFunc("/general", handlers.GetAllGeneral)
	mux.HandleFunc("/general/{id}", handlers.GetGeneralNoteById)
	mux.HandleFunc("/general/add", handlers.AddGeneralNote)
	mux.HandleFunc("/general/{id}/update", handlers.EditGeneral)
	mux.HandleFunc("/general/{id}/update/note", handlers.EditGeneralNote)
	mux.HandleFunc("/general/{id}/remove", handlers.RemoveGeneralById)

	handler := c.Handler(mux)
	err = http.ListenAndServe(":3000", handler)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

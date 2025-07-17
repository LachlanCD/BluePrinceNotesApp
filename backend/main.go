package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/handlers"
)

func main() {
	var err error

	fmt.Println("running")
	err = db_interactions.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db created")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/rooms/{workspaceID}", handlers.GetAllRooms)
	mux.HandleFunc("/api/rooms/{workspaceID}/{id}", handlers.GetRoomById)
	mux.HandleFunc("/api/rooms/{workspaceID}/add", handlers.AddRoom)
	mux.HandleFunc("/api/rooms/{workspaceID}/{id}/update", handlers.EditRoom)
	mux.HandleFunc("/api/rooms/{workspaceID}/{id}/update/note", handlers.EditRoomNote)
	mux.HandleFunc("/api/rooms/{workspaceID}/{id}/remove", handlers.RemoveRoomById)

	mux.HandleFunc("/api/general/{workspaceID}", handlers.GetAllGeneral)
	mux.HandleFunc("/api/general/{workspaceID}/{id}", handlers.GetGeneralNoteById)
	mux.HandleFunc("/api/general/{workspaceID}/add", handlers.AddGeneralNote)
	mux.HandleFunc("/api/general/{workspaceID}/{id}/update", handlers.EditGeneral)
	mux.HandleFunc("/api/general/{workspaceID}/{id}/update/note", handlers.EditGeneralNote)
	mux.HandleFunc("/api/general/{workspaceID}/{id}/remove", handlers.RemoveGeneralById)

	mux.HandleFunc("/api/create-workspace", handlers.CreateWorkspaceID)

	distDir := "./frontend/dist"
	fs := http.FileServer(http.Dir(distDir))

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If path starts with /api/, delegate to API mux
		if strings.HasPrefix(r.URL.Path, "/api/") {
			mux.ServeHTTP(w, r)
			return
		}

		// Try to serve static files
		path := filepath.Join(distDir, r.URL.Path)
		info, err := os.Stat(path)

		// If file exists and is not a directory
		if err == nil && !info.IsDir() {
			fs.ServeHTTP(w, r)
			return
		}

		// Fallback to index.html (SPA)
		http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
	})

	err = http.ListenAndServe(":4000", rootHandler)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

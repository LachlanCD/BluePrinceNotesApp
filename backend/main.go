package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/handlers"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rooms", handlers.GetRooms)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

package handlers

import (
	"fmt"
	"net/http"
)


func printRequest(r *http.Request) {
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL.String())
}

package handlers

import (
	"database/sql"
	"io"
	"net/http"
	"testing"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func openDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	return db.Ping()
}

func initTestingDB() {
	db_interactions.InitDB("../../data/test.db")
	openDB("../../data/test.db")
	seedDB()
}

func seedDB() {
	rooms := []models.Room{
		{
			Id:     1,
			Name:   "room1",
			Colour: "col1",
			Notes:  "note1",
		},
		{
			Id:     2,
			Name:   "room2",
			Colour: "col2",
			Notes:  "note2",
		},
		{
			Id:     3,
			Name:   "room3",
			Colour: "col3",
			Notes:  "note3",
		},
	}

	generals := []models.General{
		{
			Id:    1,
			Name:  "gen1",
			Notes: "note1",
		},
		{
			Id:    2,
			Name:  "gen2",
			Notes: "note2",
		},
	}

	for _, i := range rooms {
		query := "INSERT INTO rooms (id, name, colour, notes) VALUES (?,?,?,?)"
		db.Exec(query, i.Id, i.Name, i.Colour, i.Notes)
	}
	for _, i := range generals {
		query := "INSERT INTO general (id, name, notes) VALUES (?,?,?)"
		db.Exec(query, i.Id, i.Name, i.Notes)
	}

}

func cleanDB() {
	query := "DROP TABLE rooms"
	db.Exec(query)
	query = "DROP TABLE general"
	db.Exec(query)
}

func getBody(res *http.Response, t *testing.T) []byte {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	defer res.Body.Close()
	return body
}

func checkStatus(expected int, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Status Code\ngot %d\n want %d", actual, expected)
	}
}

func checkBody(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Body\ngot %q\n want %q", actual, expected)
	}
}

package db_interactions

import (
	"database/sql"

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

func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS Rooms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			workspace_id TEXT NOT NULL,
			name TEXT NOT NULL,
			colour TEXT,
			notes TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS General (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			workspace_id TEXT NOT NULL,
			name TEXT NOT NULL,
			notes TEXT
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func readAllRooms(workspaceID string) ([]*models.Room, error) {
	query := "SELECT id, name, colour FROM rooms WHERE workspace_id=?"
	rows, err := db.Query(query, workspaceID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rooms []*models.Room

	for rows.Next() {
		room := models.Room{}
		if err := rows.Scan(&room.Id, &room.Name, &room.Colour); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func readRoomById(workspaceID string, id int) (*models.Room, error) {
	query := "SELECT id, name, colour, notes FROM rooms WHERE id=? AND workspace_id=?"
	row := db.QueryRow(query, id, workspaceID)

	room := &models.Room{}

	if err := row.Scan(&room.Id, &room.Name, &room.Colour, &room.Notes); err != nil {
		return nil, err
	}

	return room, nil
}

func readAllGeneral(workspaceID string) ([]*models.General, error) {
	query := "SELECT id, name FROM general WHERE workspace_id=?"
	rows, err := db.Query(query, workspaceID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var generalNotes []*models.General

	for rows.Next() {
		general := models.General{}
		if err := rows.Scan(&general.Id, &general.Name); err != nil {
			return nil, err
		}
		generalNotes = append(generalNotes, &general)
	}
	return generalNotes, nil
}

func readGeneralById(workspaceID string, id int) (*models.General, error) {
	query := "SELECT id, name, notes FROM general Where id=? AND workspace_id=?"
	row := db.QueryRow(query, id, workspaceID)

	generalNote := &models.General{}

	if err := row.Scan(&generalNote.Id, &generalNote.Name, &generalNote.Notes); err != nil {
		return nil, err
	}

	return generalNote, nil
}

func getLastId(result sql.Result, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	roomid := int(lastid)
	return roomid, nil
}

func addRoom(workspaceID string, room models.Room) (int, error) {
	result, err := db.Exec("INSERT INTO rooms (workspace_id, name, colour, notes) VALUES (?, ?, ?, ?)", workspaceID, room.Name, room.Colour, "")
	return getLastId(result, err)
}

func addGeneral(workspaceID string, general models.General) (int, error) {
	result, err := db.Exec("INSERT INTO general (workspace_id, name, notes) VALUES (?, ?, ?)", workspaceID, general.Name, "")
	return getLastId(result, err)
}

func updateRoom(workspaceID string, room models.Room) error {
	query := "UPDATE rooms SET name=?, colour=? WHERE id=? AND workspace_id=?"
	_, err := db.Exec(query, room.Name, room.Colour, room.Id, workspaceID)
	return err
}

func updateRoomNote(workspaceID string, room models.Room) error {
	query := "UPDATE rooms SET notes=? WHERE id=? AND workspace_id=?"
	_, err := db.Exec(query, room.Notes, room.Id, workspaceID)
	return err
}

func updateGeneral(workspaceID string, generalNote models.General) error {
	query := "UPDATE general SET name=? WHERE id=? AND workspace_id=?"
	_, err := db.Exec(query, generalNote.Name, generalNote.Id, workspaceID)
	return err
}

func updateGeneralNote(workspaceID string, generalNote models.General) error {
	query := "UPDATE general SET notes=? WHERE id=? AND workspace_id=?"
	_, err := db.Exec(query, generalNote.Notes, generalNote.Id, workspaceID)
	return err
}

func removeRoomEntry(workspaceID string, id int) error {
	query := "DELETE FROM rooms WHERE id=? AND workspace_id=?"
	_, err := db.Exec(query, id, workspaceID)
	return err
}

func removeGeneralEntry(workspaceID string, id int) error {
	query := "DELETE FROM general WHERE id=? AND workspace_id=?"
	_, err := db.Exec(query, id, workspaceID)
	return err
}

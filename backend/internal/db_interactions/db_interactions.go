package db_interactions

import (
	"database/sql"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func openDB(dbPath string) (error) {
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
			name TEXT NOT NULL UNIQUE,
			colour TEXT,
			notes TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS General (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
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

func readAllRooms() ([]*models.Room, error) {
	query := "SELECT id, name, colour FROM rooms"
	rows, err := db.Query(query)
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

func readRoomById(id int) (*models.Room, error) {
	query := "SELECT id, name, colour, notes FROM rooms Where id=?"
	row := db.QueryRow(query, id)

	room := &models.Room{}

	if err := row.Scan(&room.Id, &room.Name, &room.Colour, &room.Notes); err != nil {
		return nil, err
	}

	return room, nil
}

func readAllGeneral() ([]*models.General, error) {
	query := "SELECT id, name FROM general"
	rows, err := db.Query(query)
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

func readGeneralById(id int) (*models.General, error) {
	query := "SELECT id, name, notes FROM general Where id=?"
	row := db.QueryRow(query, id)

	generalNote := &models.General{}

	if err := row.Scan(&generalNote.Id, &generalNote.Name, &generalNote.Notes); err != nil {
		return nil, err
	}

	return generalNote, nil
}

func getLastId(result sql.Result, err error) (int, error){
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

func addRoom(room models.Room) (int, error) {
	result, err := db.Exec("INSERT INTO rooms (name, colour, notes) VALUES (?, ?, ?)", room.Name, room.Colour, "")
	return getLastId(result, err)
}

func addGeneral(general models.General) (int, error) {
	result, err := db.Exec("INSERT INTO general (name, notes) VALUES (?, ?)", general.Name, "")
	return getLastId(result, err)
}

func updateRoom(room models.Room) (error) {
	query := "UPDATE rooms SET name=?, colour=?, notes=? WHERE id=?"
	_, err := db.Exec(query, room.Name, room.Colour, room.Notes, room.Id)
	return err
}

func updateGeneralNote(generalNote models.General) (error) {
	query := "UPDATE general SET name=?, notes=? WHERE id=?"
	_, err := db.Exec(query, generalNote.Name, generalNote.Notes, generalNote.Id)
	return err
}

func removeRoomEntry(id int) (error) {
	query := "DELETE FROM rooms WHERE id=?"
	_, err := db.Exec(query, id)
	return err
}

func removeGeneralEntry(id int) (error) {
	query := "DELETE FROM general WHERE id=?"
	_, err := db.Exec(query, id)
	return err
}

package db_interactions

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func openDB() error {
	server := os.Getenv("DB_SERVER")
	port := 1433
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	fmt.Printf(`server: %s`, server)
	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=true",
		server, user, password, port, database)

	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}

	return db.Ping()
}

func createTables() error {
	queries := []string{
		`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Rooms' AND xtype='U')
			BEGIN
					CREATE TABLE Rooms (
							id INT IDENTITY(1,1) PRIMARY KEY,
							workspace_id NVARCHAR(255) NOT NULL,
							name NVARCHAR(255) NOT NULL,
							colour NVARCHAR(255),
							notes NVARCHAR(MAX)
					);
			END`,
		`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='General' AND xtype='U')
			BEGIN
					CREATE TABLE General (
							id INT IDENTITY(1,1) PRIMARY KEY,
							workspace_id NVARCHAR(255) NOT NULL,
							name NVARCHAR(255) NOT NULL,
							notes NVARCHAR(MAX)
					);
			END`,
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
	query := "SELECT id, name, colour FROM rooms WHERE workspace_id = @p1"
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
	query := "SELECT id, name, colour, notes FROM rooms WHERE id = @p1 AND workspace_id = @p2"
	row := db.QueryRow(query, id, workspaceID)

	room := &models.Room{}

	if err := row.Scan(&room.Id, &room.Name, &room.Colour, &room.Notes); err != nil {
		return nil, err
	}

	return room, nil
}

func readAllGeneral(workspaceID string) ([]*models.General, error) {
	query := "SELECT id, name FROM general WHERE workspace_id=@p1"
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
	query := "SELECT id, name, notes FROM general Where id=@p1 AND workspace_id=@p2"
	row := db.QueryRow(query, id, workspaceID)

	generalNote := &models.General{}

	if err := row.Scan(&generalNote.Id, &generalNote.Name, &generalNote.Notes); err != nil {
		return nil, err
	}

	return generalNote, nil
}

func addRoom(workspaceID string, room models.Room) (int, error) {
	_, err := db.Exec("INSERT INTO rooms (workspace_id, name, colour, notes) VALUES (@p1, @p2, @p3, @p4)", workspaceID, room.Name, room.Colour, "")
	return 1, err
}

func addGeneral(workspaceID string, general models.General) (int, error) {
	_, err := db.Exec("INSERT INTO general (workspace_id, name, notes) VALUES (@p1, @p2, @p3)", workspaceID, general.Name, "")
	return 1, err
}

func updateRoom(workspaceID string, room models.Room) error {
	query := "UPDATE rooms SET name=@p1, colour=@p2 WHERE id=@p3 AND workspace_id=@p3"
	_, err := db.Exec(query, room.Name, room.Colour, room.Id, workspaceID)
	return err
}

func updateRoomNote(workspaceID string, room models.Room) error {
	query := "UPDATE rooms SET notes=@p1 WHERE id=@p2 AND workspace_id=@p3"
	_, err := db.Exec(query, room.Notes, room.Id, workspaceID)
	return err
}

func updateGeneral(workspaceID string, generalNote models.General) error {
	query := "UPDATE general SET name=@p1 WHERE id=@p2 AND workspace_id=@p3"
	_, err := db.Exec(query, generalNote.Name, generalNote.Id, workspaceID)
	return err
}

func updateGeneralNote(workspaceID string, generalNote models.General) error {
	query := "UPDATE general SET notes=@p1 WHERE id=@p2 AND workspace_id=@p3"
	_, err := db.Exec(query, generalNote.Notes, generalNote.Id, workspaceID)
	return err
}

func removeRoomEntry(workspaceID string, id int) error {
	query := "DELETE FROM rooms WHERE id=@p1 AND workspace_id=@p2"
	_, err := db.Exec(query, id, workspaceID)
	return err
}

func removeGeneralEntry(workspaceID string, id int) error {
	query := "DELETE FROM general WHERE id=@p1 AND workspace_id=@p2"
	_, err := db.Exec(query, id, workspaceID)
	return err
}

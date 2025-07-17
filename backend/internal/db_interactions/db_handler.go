package db_interactions

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
)

func InitDB() error {

	err := openDB()
	if err != nil {
		return err
	}

	return createTables()

}

func ReadAllRooms(workspaceID string) ([]*models.Room, error) {
	return readAllRooms(workspaceID)
}

func ReadRoomById(workspaceID string, id int) (*models.Room, error) {
	return readRoomById(workspaceID, id)
}

func ReadAllGeneral(workspaceID string) ([]*models.General, error) {
	return readAllGeneral(workspaceID)
}

func ReadGeneralById(workspaceID string, id int) (*models.General, error) {
	return readGeneralById(workspaceID, id)
}

func AddRoom(workspaceID string, room models.Room) (int, error) {
	return addRoom(workspaceID, room)
}

func AddGeneral(workspaceID string, generalNote models.General) (int, error) {
	return addGeneral(workspaceID, generalNote)
}

func RemoveRoomNote(workspaceID string, id int) error {
	return removeRoomEntry(workspaceID, id)
}

func RemoveGeneralNote(workspaceID string, id int) error {
	return removeGeneralEntry(workspaceID, id)
}

func UpdateRoom(workspaceID string, room models.Room) error {
	oldRoom, err := readRoomById(workspaceID, room.Id)
	if err != nil {
		return err
	}
	if oldRoom.Name == room.Name && oldRoom.Colour == room.Colour {
		return errors.New("Room must be updated")
	}
	return updateRoom(workspaceID, room)
}

func UpdateRoomNote(workspaceID string, room models.Room) error {
	oldRoom, err := readRoomById(workspaceID, room.Id)
	if err != nil {
		return err
	}
	if oldRoom.Notes == room.Notes {
		return errors.New("Room must be updated")
	}
	return updateRoomNote(workspaceID, room)
}

func UpdateGeneral(workspaceID string, generalNote models.General) error {
	oldGen, err := readGeneralById(workspaceID, generalNote.Id)
	if err != nil {
		return err
	}
	if oldGen.Name == generalNote.Name {
		return errors.New("General Note must be updated")
	}
	return updateGeneral(workspaceID, generalNote)
}

func UpdateGeneralNote(workspaceID string, generalNote models.General) error {
	oldGen, err := readGeneralById(workspaceID, generalNote.Id)
	if err != nil {
		return err
	}
	if oldGen.Notes == generalNote.Notes {
		return errors.New("Room must be updated")
	}
	return updateGeneralNote(workspaceID, generalNote)
}

func GenerateWorkspaceID() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}	

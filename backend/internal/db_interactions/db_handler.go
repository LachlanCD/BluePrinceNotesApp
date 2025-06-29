package db_interactions

import (
	"errors"
	"github.com/LachlanCD/BluePrinceNotesApp/internal/models"
)

func InitDB(dbPath string) error {

	err := openDB(dbPath)
	if err != nil {
		return err
	}

	return createTables()

}

func ReadAllRooms() ([]*models.Room, error) {
	return readAllRooms()
}

func ReadRoomById(id int) (*models.Room, error) {
	return readRoomById(id)
}

func ReadAllGeneral() ([]*models.General, error) {
	return readAllGeneral()
}

func ReadGeneralById(id int) (*models.General, error) {
	return readGeneralById(id)
}

func AddRoom(room models.Room) (int, error) {
	return addRoom(room)
}

func AddGeneral(generalNote models.General) (int, error) {
	return addGeneral(generalNote)
}

func RemoveRoomNote(id int) error {
	return removeRoomEntry(id)
}

func RemoveGeneralNote(id int) error {
	return removeGeneralEntry(id)
}

func UpdateRoom(room models.Room) error {
	oldRoom, err := readRoomById(room.Id)
	if err != nil {
		return err
	}
	if oldRoom == &room {
		return errors.New("Room must be updated")
	}
	return updateRoom(room)
}

func UpdateGeneral(generalNote models.General) error {
	oldGen, err := readGeneralById(generalNote.Id)
	if err != nil {
		return err
	}
	if oldGen == &generalNote {
		return errors.New("General Note must be updated")
	}
	return updateGeneralNote(generalNote)
}

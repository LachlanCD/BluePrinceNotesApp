package handlers

import "github.com/LachlanCD/BluePrinceNotesApp/internal/db_interactions"

func initTestingDB() {
	db_interactions.InitDB("../../data/test.db")
}

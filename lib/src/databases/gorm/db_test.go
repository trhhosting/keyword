package gorm

import "testing"

func TestCreateDatabaseTables(t *testing.T) {
	DropDb()
	CreateDb()
	err := CreateDatabaseTables()
	if err != nil {
		t.Fail()
	}
}

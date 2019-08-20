package gorm

import "testing"

func TestDropDb(t *testing.T) {
	DropDb()
}

func TestCreateDb(t *testing.T) {
	CreateDb()
	//grantPermissionDatabase()
}

package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"path/filepath"
	"sc.tpnfc.us/tricll/triclldb/lib/src/git"
	"time"
)

var (
	db *leveldb.DB
)

// TLMDataStore is the data store for tlm local configurations and tracking of certificate issues and renewals
type TLMDataStore struct {
	Data data
	Path string
}

// Create
func (t TLMDataStore) Create() error {
	/*
	 */
	openDataStore()
	defer closeDataStore()
	db.Put([]byte("init"), []byte(time.Now().String()), nil)
	return nil

}

// Update
func (t TLMDataStore) Update() error {
	/*
		Updates data in the store
	*/
	openDataStore()
	defer closeDataStore()

	err := db.Delete(t.Data.Key, nil)
	err = db.Put(t.Data.Key, t.Data.Data, nil)
	if err != nil {
		log.Printf("Update failed :%s\n", err)
	}
	return err
}

// Save
func (t TLMDataStore) Save() error {
	/*

	 */

	openDataStore()
	defer closeDataStore()

	err := db.Put(t.Data.Key, t.Data.Data, nil)
	if err != nil {
		log.Printf("Save failed :%s\n", err)
	}
	return err
}

// Delete
func (t TLMDataStore) Delete() error {
	/*

	 */
	openDataStore()
	defer closeDataStore()

	err := db.Delete(t.Data.Key, nil)
	if err != nil {
		log.Printf("Delete failed :%s\n", err)
	}
	return err
}

// Lookup
func (t TLMDataStore) Lookup() ([]byte, error) {
	openDataStore()
	defer closeDataStore()

	val, err := db.Get(t.Data.Key, nil)
	return val, err
}
func (t TLMDataStore) Loop() ([][]byte, error) {
	//TODO implement loop over data
	return nil, nil
}

// used by TLMDataStore don't delete !!!!
func openDataStore() error {
	var err error
	path := filepath.Join(git.HomeDir, "tridata")
	db, err = leveldb.OpenFile(path, nil)
	return err
}
func closeDataStore() {
	db.Close()
}

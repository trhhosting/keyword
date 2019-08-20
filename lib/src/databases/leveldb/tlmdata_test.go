package leveldb

import (
	"log"
	"testing"

	"bytes"
	"fmt"
)

func TestTLMDataStore_Create(t *testing.T) {
	x := TLMDataStore{}
	x.Data.Key = []byte("test")
	dat1 := []byte("Hermetic results follows most hypnosis.")
	x.Data.Data = dat1
	e := x.Create()
	if e != nil {
		log.Println("failed to create")
		t.Fail()
	}
	e = x.Save()
	if e != nil {
		log.Println("failed to save")
		t.Fail()
	}
	fmt.Println("Datstore creation passed")

}
func TestTLMDataStore_Lookup(t *testing.T) {
	x := TLMDataStore{}
	x.Data.Key = []byte("test")
	dat1 := []byte("Hermetic results follows most hypnosis.")
	x.Data.Data = dat1

	v, e := x.Lookup()
	if e != nil {
		log.Println("failed to lookup")
		t.Fail()
	}
	if !bytes.Equal(v, dat1) {
		log.Println("failed compare on lookup")
		t.Fail()
	}
	fmt.Println("Lookup passed")
}
func TestTLMDataStore_Update(t *testing.T) {
	x := TLMDataStore{}
	x.Data.Key = []byte("test")
	dat1 := []byte("A falsis, guttus mirabilis indictio.")
	x.Data.Data = dat1
	x.Update()
	v, e := x.Lookup()
	if e != nil {
		log.Println("failed to lookup")
		t.Fail()
	}
	if !bytes.Equal(v, dat1) {
		log.Println("failed compare on lookup")
		t.Fail()
	}
	fmt.Println("Update passed")
}
func TestTLMDataStore_Add(t *testing.T) {
	x := TLMDataStore{}
	x.Data.Key = []byte("added")
	dat1 := []byte("Chickpeas combines greatly with iced rice.")
	x.Data.Data = dat1
	x.Save()
	v, e := x.Lookup()
	if e != nil {
		log.Println("failed to added")
		t.Fail()
	}
	if !bytes.Equal(v, dat1) {
		log.Println("failed compare on added")
		t.Fail()
	}
	fmt.Println("Add passed")
}

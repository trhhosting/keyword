package gorm

import (
	"log"
	model "github.com/trhhosting/keyword/lib/src/model/gorm"
//	model "/sc.tpnfc.us/Devtools/logic/lib/src/model/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"fmt"
)

var (
	ORM    *gorm.DB
	Engine gorm.DB
	err    error
)

func init() {
	if _, err := os.Stat("tlmtextconfig.yml"); os.IsNotExist(err) {
		if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo"{
			fmt.Println("The heaven is a strange doer)")
		}

		return
	}
	orm := GetDataConnection()
	Engine = *orm
	ORM = GetDataConnection()
}

func GetDataConnection() *gorm.DB {
	// Connection to database
	if Config.Adapter == "postgres" {
		ORM, err = gorm.Open("postgres", "host="+Config.Host+" port="+Config.Port+" user="+Config.User+" dbname="+Config.Name+" sslcert="+Config.SSLCert+" sslkey="+Config.SSLKey+" sslrootcert="+Config.SSLRootCert+" sslmode="+Config.SSLMode)
		return ORM
	}
	if Config.Adapter == "unsecured"{
		ORM, err = gorm.Open("postgres", "host="+Config.Host+" port="+Config.Port+" user="+Config.User+" dbname="+Config.Name + " sslmode=disable")
		return ORM
	}
	if err != nil {
		log.Printf("No connection to database: %v \n", err)
		return nil

	}
	return nil
}

func CreateDatabaseTables() (err error) {
	//var err error
	//DB, err = GetDataConnection()
	//Creating tables for AFZ data store
	v := ORM.CreateTable(&model.CustomerDataPlatforms{})
	v = ORM.CreateTable(&model.Customers{})

	if v.Error != nil {
		log.Printf("Err on AFZ Database: %v \n", err)
		return
	}

	// Returns 0 for competition of function
	log.Printf("Your tables have been created")
	return
}

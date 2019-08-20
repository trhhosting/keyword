package gorm

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"github.com/trhhosting/keyword/lib/src/config"
	"time"
	"os"
)

var (
	OrmRoot *sql.DB
	Config  = config.CockroachConfig()
)

func init() {
	if _, err := os.Stat("tlmtextconfig.yml"); os.IsNotExist(err) {
		if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo"{
			fmt.Printf("TLMConfig-init--> xdat\n")
			fmt.Println("Must run init first !! :)")
		}

		return
	}

	orm, er := getDataConnectionRoot()
	if er != nil {
		log.Printf("Connection failed : %v \n", er)
	}
	OrmRoot = orm
}

/*	drop DATABASE IF EXISTS legion_market;
	CREATE DATABASE legion_market;
	grant ALL on DATABASE legion_market to aqualobster;
*/

func createDb() (err error) {
	x, err := OrmRoot.Exec("CREATE DATABASE " + Config.Name + ";")
	if err != nil {
		log.Fatal(err)
	}
	grantPermissionDatabase()
	log.Printf("Database %v has been created with %v \n", Config.Name, x)
	return
}
func dropDB(f bool) (int8, error) {
	// If true force delete already created database

	if f == true {
		x, er := OrmRoot.Exec("drop DATABASE " + Config.Name + ";")
		if er != nil {
			fmt.Printf("Was not able to drop Database \n")
		}
		log.Printf("Database %v has been DROPPED with %v \n", Config.Name, x)
		return 7, er
	} else {
		x, er := OrmRoot.Exec("drop DATABASE IF EXISTS " + Config.Name + ";")
		if er != nil {
			fmt.Printf("Wasn't  able to drop Database, database might already exsist \n")
		}
		log.Printf("Database %v has been DROPPED with %v \n", Config.Name, x)
		return 8, er
	}
	er := errors.New("did not drop database")
	return 3, er
}
func grantPermissionDatabase() (int8, error) {
	createDatabaseUser()
	x, er := OrmRoot.Exec("GRANT ALL on DATABASE " + Config.Name + " TO " + Config.User + ";")

	if er != nil {
		er = errors.New("was not able to grant access")
	}
	log.Printf("User: %v has been granted access to %v with %v \n", Config.User, Config.Name, x)

	defer OrmRoot.Close()
	return 0, er
}
func createDatabaseUser() (int8, error) {
	x, er := OrmRoot.Exec("CREATE USER " + Config.User + ";")
	if er != nil {
		er = errors.New("was not able to create user")
		return 1, er
	}
	log.Printf("User: %v has been created %v \n", Config.User,  x)
	return 0, er
}

func getDataConnectionRoot() (OrmRoot *sql.DB, er error) {
	if Config.Adapter == "postgres" {
		OrmRoot, er = sql.Open("postgres", "host="+Config.Host+" port="+Config.Port+" user="+Config.Root+" dbname="+Config.Name+" sslcert="+Config.RootCert+" sslkey="+Config.RootKey+" sslrootcert="+Config.SSLRootCert+" sslmode="+Config.SSLMode)
		return
	} else if Config.Adapter == "unsecured"{
		OrmRoot, er = sql.Open("postgres", "host="+Config.Host+" port="+Config.Port+" user="+Config.Root + " sslmode=disable")
		fmt.Println(OrmRoot)
		if er != nil {
		   log.Fatal(er)
		}
		return
	} else {
		panic(errors.New("not supported database adapter \n"))
	}
	return nil, nil
}

func CreateDb() {
		x := createDb()
		if x == nil {
			fmt.Printf("Database created %v \n", time.Now())
		}
		grantPermissionDatabase()
		log.Printf("Just saying")
}
func DropDb() {

	x, er := dropDB(true)
	if er != nil {
		log.Printf("Database not dropped %v \n", er)
	}
	if x == 3 {
		log.Printf("Database not dropped \n")
	}
	if x == 7 || x == 8 {
		log.Printf("Database dropped \n")
	}
}

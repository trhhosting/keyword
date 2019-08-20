package gorm

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetDataSQLConnection() (db *sql.DB) {

	db, err = sql.Open("postgres", "host="+Config.Host+" port="+Config.Port+" user="+Config.User+""+
		" dbname="+Config.Name+" sslcert="+Config.SSLCert+" sslkey="+Config.SSLKey+
		" sslrootcert="+Config.SSLRootCert+" sslmode="+Config.SSLMode)
	if err != nil {
		panic(err)
	}
	return
}

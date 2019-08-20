package config

import (
	"log"
	"os"
	"path/filepath"
	"sc.tpnfc.us/oyoshi/configor"
)

type Natscluster struct {
	Certfile string
	Keyfile  string
	Cacert   string
	Address  string
	Port     string
}

type Cockroachdatastore struct {
	Name        string
	Adapter     string
	Host        string
	Port        string
	User        string
	Password    string
	SSLMode     string
	SSLCert     string
	SSLKey      string
	SSLRootCert string
	Root        string
	RootCert    string
	RootKey     string
}

var TLMNATsConfig = struct {
	Natscluster struct {
		Certfile string
		Keyfile  string
		Cacert   string
		Address  string
		Port     string
	}
	Cockroachdatastore struct{
		Name        string
		Adapter     string
		Host        string
		Port        string
		User        string
		Password    string
		SSLMode     string
		SSLCert     string
		SSLKey      string
		SSLRootCert string
		Root        string
		RootCert    string
		RootKey     string
	}
}{}

func loadFiles() {
	path, _ := os.Getwd()

	x := filepath.Join(path, "tlmtextconfig.yml")


	er := configor.Load(&TLMNATsConfig, x)
	if er != nil {
	   log.Fatal(er)
	}
	if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo" {
		log.Printf("Filelocation: %v \n", x)
	}


}

func CockroachConfig() (c Cockroachdatastore) {
	loadFiles()
	c = TLMNATsConfig.Cockroachdatastore
	return
}
func NatsClusterConfig() (n Natscluster) {
	loadFiles()
	n = TLMNATsConfig.Natscluster
	return
}

package config

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TestTLMConfigData(t *testing.T) {
	p, _ := os.Getwd()
	path := filepath.Join(p, "out")
	cockroachcerts := filepath.Join(p, "cockroachCerts")
	h, _ := os.Hostname()
	keyPath := filepath.Join(path, h)
	caPath := filepath.Join(path, "nats-ca.crt")
	var x = TLMConfigData{
		CertFile: keyPath + ".crt",
		KeyFile:  keyPath + ".key",
		CACert:   caPath,
		Address:  "tls://192.168.1.140",
		Port:     "443",
		Name: "activesms",
		Host: "192.168.1.247",
		CockroachPort: "26257",
		User: "intensiveplastic",
		SSLMode: "verify-ca",
		SSLCert: cockroachcerts + "/client.intensiveplastic.crt",
		SSLKey: cockroachcerts + "/client.intensiveplastic.key",
		SSLRootcert: cockroachcerts + "/ca.crt",
		Rootcert: cockroachcerts + "/client.root.crt",
		Rootkey: cockroachcerts + "/client.root.key",
	}
	dat ,e := x.CreateConfig()
	if e != nil {
		t.Fail()
	}
	if dat == ""{
		t.Fail()
	}
	e = ioutil.WriteFile("tlmtextconfig.yml", []byte(dat), 0744)
	if e != nil {
		t.Fail()
	}
}
func TestCockroachConfig(t *testing.T) {
	c := CockroachConfig()
	if c.User != "intensiveplastic"{
		t.Fail()
	}
	if c.Port != "26257" {
		t.Fail()
	}
	//fmt.Println(c)
}
func TestNatsClusterConfig(t *testing.T) {
	n := NatsClusterConfig()
	if n.Certfile == ""{
		t.Fail()
	}
}
func TestLoadFile(t *testing.T) {
	loadFiles()
}
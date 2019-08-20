package config

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TestTLMConfigDataOnly(t *testing.T) {
	p, _ := os.Getwd()
	path := filepath.Join(p, "out")
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
		SSLCert: "cockroachCerts/client.intensiveplastic.crt",
		SSLKey: "cockroachCerts/client.intensiveplastic.key",
		SSLRootcert: "cockroachCerts/ca.crt",
		Rootcert: "cockroachCerts/client.root.crt",
		Rootkey: "cockroachCerts/client.root.key",
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
//URL:      "https://natsdeploy:VNAuGu5b9qgMdUNqVBxl@sc.tpnfc.us/NATs/serverlists/raw/master/local.list",

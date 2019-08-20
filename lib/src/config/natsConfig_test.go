package config

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TestNATsConfigDataOnly(t *testing.T) {
	p, _ := os.Getwd()
	path := filepath.Join(p, "out")
	h, _ := os.Hostname()
	keyPath := filepath.Join(path, h)
	caPath := filepath.Join(path, "nats-ca.crt")

	routes :="\t nats-route://192.168.1.140:8443 \n\t nats-route://192.168.1.156:8443"
	var x = NatsServerConfigData{
		Listen: "0.0.0.0:443",
		CertFile: keyPath + ".crt",
		KeyFile: keyPath + ".key",
		CACert: caPath,
		Host: "0.0.0.0",
		HostPort: "8443",
		NatsRoute: routes,
	}
	dat ,e := x.CreateConfig()
	if e != nil {
		t.Fail()
	}
	if dat == ""{
		t.Fail()
	}
	e = ioutil.WriteFile("natsserver.conf", []byte(dat), 0744)
	if e != nil {
		t.Fail()
	}
}

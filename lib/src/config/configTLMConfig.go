package config

import (
	"github.com/SpivEgin/fasttemplate"
	"strings"
)

const tlmconf = `---
natscluster:
  certfile: "{&certfile&}"
  keyfile: "{&keyfile&}"
  cacert: "{&cacert&}"
  address: "{&address&}"
  port: "{&port&}"
cockroachdatastore:
  adapter: "postgres"
  name: "{&name&}"
  host: "{&host&}"
  port: {&cockroachport&}
  user: "{&user&}"
  password: ""
  sslmode: "verify-ca"
  sslcert: "{&sslcert&}"
  sslkey: "{&sslkey&}"
  sslrootcert: "{&sslrootcert&}"
  root: "root"
  rootcert: "{&rootcert&}"
  rootkey: "{&rootkey&}"

...
`

//TLMConfigData used to write config file after the certs are signed
type TLMConfigData struct {
	CertFile      string
	KeyFile       string
	CACert        string
	Address       string
	Port          string
	Name          string
	Host          string
	CockroachPort string
	User          string
	SSLMode       string
	SSLCert       string
	SSLKey        string
	SSLRootcert   string
	Root          string
	Rootcert      string
	Rootkey       string
}

func (c TLMConfigData) CreateConfig() (config string, err error) {
	t, err := fasttemplate.NewTemplate(tlmconf, "{&", "&}")

	config = t.ExecuteString(map[string]interface{}{
		"certfile":      convertPaths(c.CertFile),
		"keyfile":       convertPaths(c.KeyFile),
		"cacert":        convertPaths(c.CACert),
		"address":       c.Address,
		"port":          c.Port,
		"cockroachport": c.CockroachPort,
		"name":          c.Name,
		"host":          c.Host,
		"CockroachPort": c.CockroachPort,
		"user":          c.User,
		"sslcert":       convertPaths(c.SSLCert),
		"sslkey":        convertPaths(c.SSLKey),
		"sslrootcert":   convertPaths(c.SSLRootcert),
		"rootcert":      convertPaths(c.Rootcert),
		"rootkey":       convertPaths(c.Rootkey),
	})
	return
}
func convertPaths(filepath string) (file string) {
	filepath = strings.Replace(filepath, "C:", "", 1)
	file = strings.Replace(filepath, `\`, `/`, 100)

	return
}

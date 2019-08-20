package config

import (
	"github.com/SpivEgin/fasttemplate"
	"fmt"
)

const natsserverconf =
	`# TLM NATs TLS config file

listen: {&listen&}

tls {
  cert_file:  "{&certfile&}"
  key_file:   "{&keyfile&}"
  ca_file:   "{&cacert&}"
  verify:    true
  timeout:    15
}

cluster {
  host: {&host&}
  port: {&hostport&}

  routes = [
   {&natsroute&}
  ]
}
`
const NATSDConf =
	`# TLM NATs TLS config file

listen: {&listen&}

tls {
  cert_file:  "{&certfile&}"
  key_file:   "{&keyfile&}"
  ca_file:   "{&cacert&}"
  verify:    true
  timeout:    15
}

cluster {
  host: {&host&}
  port: {&hostport&}
  ]
}
`
//NatsServerConfigData used to write config file for nats Server config
type NatsServerConfigData struct {
	Listen string
	CertFile string
	KeyFile  string
	CACert   string
	Host  string
	HostPort     string
	NatsRoute string

}

func (n NatsServerConfigData) CreateConfig()  (config string, err error) {
	t, err := fasttemplate.NewTemplate(natsserverconf, "{&", "&}")

	config = t.ExecuteString(map[string]interface{}{
		"listen": n.Listen,
		"certfile": convertPaths(n.CertFile),
		"keyfile": convertPaths(n.KeyFile),
		"cacert": convertPaths(n.CACert),
		"host": n.Host,
		"hostport": n.HostPort,
		"natsroute": n.NatsRoute,
	})

	return
}

func (n NatsServerConfigData) NATSDConfig()  (config string, err error) {
	t, err := fasttemplate.NewTemplate(NATSDConf, "{&", "&}")

	config = t.ExecuteString(map[string]interface{}{
		"listen": n.Listen,
		"certfile": convertPaths(n.CertFile),
		"keyfile": convertPaths(n.KeyFile),
		"cacert": convertPaths(n.CACert),
		"host": n.Host,
		"hostport": n.HostPort,
	})
	fmt.Println(config)
	return
}
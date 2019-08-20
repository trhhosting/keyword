package echoserver

import (
	"testing"
)

func TestTLMSMSServer(t *testing.T) {
	//h, _ := os.Hostname()
	var x = TLMSMSServer{
		IPAddress: "192.168.1.104",
		Port: "80",
	}
	//fmt.Println(h)
	x.TLMServer()
}

